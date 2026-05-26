package processor

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/llmx"
	"github.com/zhiguang/zhiguang-go/pkg/textx"
	"github.com/zhiguang/zhiguang-go/services/agent/indexer/internal/svc"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	counterevent "github.com/zhiguang/zhiguang-go/services/counter/shared/event"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type Processor struct {
	sc       *svc.ServiceContext
	dedup    *Dedup
	dedupTTL time.Duration
}

func New(sc *svc.ServiceContext) *Processor {
	return &Processor{
		sc:       sc,
		dedup:    NewDedup(sc.Redis),
		dedupTTL: time.Duration(sc.Config.Dedup.TtlSeconds) * time.Second,
	}
}

func (p *Processor) Handle(ctx context.Context, m kafka.Message) error {
	var ev counterevent.CounterEvent
	if err := json.Unmarshal(m.Value, &ev); err != nil {
		logx.WithContext(ctx).Errorf("agent-indexer: bad event: %v", err)
		return nil
	}
	if ev.EntityType != "knowpost" || ev.Metric != "fav" {
		return nil
	}
	postID, err := strconv.ParseInt(ev.EntityId, 10, 64)
	if err != nil || postID <= 0 || ev.UserId <= 0 {
		return nil
	}
	if ev.Delta > 0 {
		return p.upsertByFavorite(ctx, ev.UserId, postID)
	}
	if ev.Delta < 0 {
		return p.removeByUnfavorite(ctx, ev.UserId, postID)
	}
	return nil
}

func (p *Processor) upsertByFavorite(ctx context.Context, userID, postID int64) error {
	detail, err := p.sc.KnowPostRpc.GetDetail(ctx, &knowpostpb.GetDetailReq{Id: postID, ViewerId: userID})
	if err != nil {
		return err
	}
	if detail.Status != "published" {
		return nil
	}
	body, err := fetchContent(ctx, p.sc.HttpClient, detail.ContentUrl)
	if err != nil {
		return p.markFailed(ctx, userID, postID, "", err)
	}
	if strings.TrimSpace(body) == "" {
		return nil
	}

	version := hashVersion(detail.ContentSha256, body)
	dedupKey := fmt.Sprintf("dedup:agent:%d:%d:%s", userID, postID, version)
	ok, err := p.dedup.Acquire(ctx, dedupKey, p.dedupTTL)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	now := time.Now().UnixMilli()
	if _, err := p.sc.Db.ExecCtx(ctx,
		"INSERT INTO agent_knowledge_index (user_id,post_id,version_hash,status,retry_count,next_retry_at,updated_at,created_at) VALUES (?,?,?,?,0,0,?,?) ON DUPLICATE KEY UPDATE status='indexing', updated_at=VALUES(updated_at), last_error=''",
		userID, postID, version, "indexing", now, now,
	); err != nil {
		return err
	}

	sections := textx.SplitByHeader(body)
	if len(sections) == 0 {
		sections = []textx.Section{{Body: body}}
	}
	chunks := make([]chunk, 0, 64)
	pos := 0
	for _, sec := range sections {
		for _, c := range textx.Chunk(sec.Body, p.sc.Config.Chunk.Size, p.sc.Config.Chunk.Overlap) {
			pos++
			chunks = append(chunks, chunk{pos: pos, title: sec.Title, text: c})
		}
	}
	if len(chunks) == 0 {
		return p.markFailed(ctx, userID, postID, version, fmt.Errorf("no chunks"))
	}

	texts := make([]string, len(chunks))
	for i, c := range chunks {
		texts[i] = c.text
	}
	vecs, err := llmx.EmbedFloat32(ctx, p.sc.Embed, texts)
	if err != nil || len(vecs) != len(chunks) {
		if err == nil {
			err = fmt.Errorf("embedding count mismatch")
		}
		return p.markFailed(ctx, userID, postID, version, err)
	}

	if err := p.deleteOldUserPost(ctx, userID, postID); err != nil {
		return p.markFailed(ctx, userID, postID, version, err)
	}
	if err := p.deleteFromMilvus(ctx, userID, postID); err != nil {
		return p.markFailed(ctx, userID, postID, version, err)
	}

	var buf bytes.Buffer
	for i, c := range chunks {
		docID := fmt.Sprintf("u%d-p%d-c%04d", userID, postID, c.pos)
		action, _ := json.Marshal(map[string]any{"index": map[string]any{"_index": p.sc.Config.KnowledgeIndex, "_id": docID}})
		doc := map[string]any{
			"user_id":    userID,
			"post_id":    postID,
			"chunk_id":   fmt.Sprintf("c%04d", c.pos),
			"position":   c.pos,
			"title_path": strings.TrimSpace(detail.Title + " / " + c.title),
			"text":       c.text,
			"version":    version,
			"status":     "ready",
			"source_url": detail.ContentUrl,
			"updated_at": now,
			"embedding":  vecs[i],
		}
		line, _ := json.Marshal(doc)
		buf.Write(action)
		buf.WriteByte('\n')
		buf.Write(line)
		buf.WriteByte('\n')
	}
	if err := p.sc.Es.Bulk(ctx, buf.Bytes()); err != nil {
		return p.markFailed(ctx, userID, postID, version, err)
	}
	if err := p.upsertMilvusChunks(ctx, userID, postID, version, detail.Title, chunks, vecs); err != nil {
		return p.markFailed(ctx, userID, postID, version, err)
	}

	_, _ = p.sc.Db.ExecCtx(ctx,
		"UPDATE agent_knowledge_index SET status='ready', chunk_count=?, retry_count=0, next_retry_at=0, last_error='', updated_at=? WHERE user_id=? AND post_id=? AND version_hash=?",
		len(chunks), now, userID, postID, version,
	)
	logx.WithContext(ctx).Infof("agent-indexer indexed user=%d post=%d chunks=%d", userID, postID, len(chunks))
	return nil
}

func (p *Processor) upsertMilvusChunks(ctx context.Context, userID, postID int64, version, title string, chunks []chunk, vecs [][]float32) error {
	if p.sc.Milvus == nil || !p.sc.Config.Milvus.Enabled {
		return nil
	}
	ids := make([]string, 0, len(chunks))
	uids := make([]int64, 0, len(chunks))
	pids := make([]int64, 0, len(chunks))
	cids := make([]string, 0, len(chunks))
	titles := make([]string, 0, len(chunks))
	texts := make([]string, 0, len(chunks))
	versions := make([]string, 0, len(chunks))
	statuses := make([]string, 0, len(chunks))
	vectors := make([][]float32, 0, len(chunks))
	for i, c := range chunks {
		ids = append(ids, fmt.Sprintf("u%d-p%d-c%04d", userID, postID, c.pos))
		uids = append(uids, userID)
		pids = append(pids, postID)
		cids = append(cids, fmt.Sprintf("c%04d", c.pos))
		titles = append(titles, trimContent(strings.TrimSpace(title+" / "+c.title), 1024))
		texts = append(texts, trimContent(c.text, 30000))
		versions = append(versions, version)
		statuses = append(statuses, "ready")
		vectors = append(vectors, vecs[i])
	}
	_, err := p.sc.Milvus.Upsert(ctx, milvusclient.NewColumnBasedInsertOption(p.sc.Config.Milvus.Collection).
		WithVarcharColumn("id", ids).
		WithInt64Column("user_id", uids).
		WithInt64Column("post_id", pids).
		WithVarcharColumn("chunk_id", cids).
		WithVarcharColumn("title_path", titles).
		WithVarcharColumn("text", texts).
		WithVarcharColumn("version", versions).
		WithVarcharColumn("status", statuses).
		WithFloatVectorColumn(p.sc.Config.Milvus.VectorField, p.sc.Config.Milvus.VectorDim, vectors),
	)
	return err
}

func (p *Processor) removeByUnfavorite(ctx context.Context, userID, postID int64) error {
	if err := p.deleteOldUserPost(ctx, userID, postID); err != nil {
		return err
	}
	if err := p.deleteFromMilvus(ctx, userID, postID); err != nil {
		return err
	}
	_, _ = p.sc.Db.ExecCtx(ctx,
		"UPDATE agent_knowledge_index SET status='deleted', updated_at=?, last_error='' WHERE user_id=? AND post_id=?",
		time.Now().UnixMilli(), userID, postID,
	)
	logx.WithContext(ctx).Infof("agent-indexer unfavorite removed user=%d post=%d", userID, postID)
	return nil
}

func (p *Processor) deleteFromMilvus(ctx context.Context, userID, postID int64) error {
	if p.sc.Milvus == nil || !p.sc.Config.Milvus.Enabled {
		return nil
	}
	expr := fmt.Sprintf("user_id == %d and post_id == %d", userID, postID)
	_, err := p.sc.Milvus.Delete(ctx, milvusclient.NewDeleteOption(p.sc.Config.Milvus.Collection).WithExpr(expr))
	return err
}

func (p *Processor) ReconcileAndRetry(ctx context.Context, limit int) error {
	if limit <= 0 {
		limit = 200
	}
	now := time.Now().UnixMilli()
	var rows []struct {
		UserID      int64  `db:"user_id"`
		PostID      int64  `db:"post_id"`
		VersionHash string `db:"version_hash"`
		RetryCount  int    `db:"retry_count"`
	}
	err := p.sc.Db.QueryRowsCtx(ctx, &rows,
		"SELECT user_id, post_id, version_hash, retry_count FROM agent_knowledge_index WHERE status IN ('failed','indexing') AND (next_retry_at=0 OR next_retry_at<=?) ORDER BY updated_at ASC LIMIT ?",
		now, limit,
	)
	if err != nil {
		return err
	}
	for _, row := range rows {
		marked, markErr := p.sc.CounterRpc.IsMarked(ctx, &counterpb.IsMarkedReq{
			EntityType: "knowpost",
			EntityId:   strconv.FormatInt(row.PostID, 10),
			Metric:     "fav",
			UserId:     row.UserID,
		})
		if markErr != nil {
			continue
		}
		if marked == nil || !marked.Marked {
			_ = p.removeByUnfavorite(ctx, row.UserID, row.PostID)
			continue
		}
		if err := p.upsertByFavorite(ctx, row.UserID, row.PostID); err != nil {
			_ = p.bumpRetry(ctx, row.UserID, row.PostID, row.VersionHash, row.RetryCount+1, err)
		}
	}
	return nil
}

func (p *Processor) bumpRetry(ctx context.Context, userID, postID int64, version string, retry int, cause error) error {
	if retry < 1 {
		retry = 1
	}
	backoff := p.sc.Config.RetryBackoffSec * retry
	if backoff <= 0 {
		backoff = 60
	}
	if backoff > 1800 {
		backoff = 1800
	}
	nextAt := time.Now().Add(time.Duration(backoff) * time.Second).UnixMilli()
	_, err := p.sc.Db.ExecCtx(ctx,
		"UPDATE agent_knowledge_index SET status='failed', retry_count=?, next_retry_at=?, updated_at=?, last_error=? WHERE user_id=? AND post_id=? AND version_hash=?",
		retry, nextAt, time.Now().UnixMilli(), shortErr(cause), userID, postID, version,
	)
	return err
}

func (p *Processor) markFailed(ctx context.Context, userID, postID int64, version string, cause error) error {
	now := time.Now().UnixMilli()
	if strings.TrimSpace(version) == "" {
		_, _ = p.sc.Db.ExecCtx(ctx,
			"UPDATE agent_knowledge_index SET status='failed', retry_count=retry_count+1, next_retry_at=?, updated_at=?, last_error=? WHERE user_id=? AND post_id=? ORDER BY id DESC LIMIT 1",
			now+int64(p.sc.Config.RetryBackoffSec*1000), now, shortErr(cause), userID, postID,
		)
		return cause
	}
	_, _ = p.sc.Db.ExecCtx(ctx,
		"UPDATE agent_knowledge_index SET status='failed', retry_count=retry_count+1, next_retry_at=?, updated_at=?, last_error=? WHERE user_id=? AND post_id=? AND version_hash=?",
		now+int64(p.sc.Config.RetryBackoffSec*1000), now, shortErr(cause), userID, postID, version,
	)
	return cause
}

func (p *Processor) deleteOldUserPost(ctx context.Context, userID, postID int64) error {
	q := map[string]any{"query": map[string]any{"bool": map[string]any{"filter": []any{
		map[string]any{"term": map[string]any{"user_id": userID}},
		map[string]any{"term": map[string]any{"post_id": postID}},
	}}}}
	body, _ := json.Marshal(q)
	_, err := p.sc.Es.DeleteByQuery(ctx, p.sc.Config.KnowledgeIndex, body)
	return err
}

func fetchContent(ctx context.Context, hc *http.Client, url string) (string, error) {
	if strings.TrimSpace(url) == "" {
		return "", nil
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("fetch content status=%d", resp.StatusCode)
	}
	buf, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func hashVersion(contentSHA, body string) string {
	v := strings.TrimSpace(contentSHA)
	if v != "" {
		return v
	}
	s := sha256.Sum256([]byte(body))
	return hex.EncodeToString(s[:])
}

func shortErr(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	if len(s) > 240 {
		return s[:240]
	}
	return s
}

type chunk struct {
	pos   int
	title string
	text  string
}

func trimContent(s string, n int) string {
	r := []rune(strings.TrimSpace(s))
	if n <= 0 || len(r) <= n {
		return string(r)
	}
	return string(r[:n])
}
