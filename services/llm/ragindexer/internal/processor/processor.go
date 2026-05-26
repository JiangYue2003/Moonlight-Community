// Package processor rag-indexer 处理器：消费 canal-outbox（aggregate_type=knowpost）→ 切块/嵌入/写 ES 向量。
package processor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/canalx"
	"github.com/zhiguang/zhiguang-go/pkg/llmx"
	"github.com/zhiguang/zhiguang-go/pkg/textx"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	event "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
	"github.com/zhiguang/zhiguang-go/services/llm/ragindexer/internal/svc"
)

type Processor struct {
	sc       *svc.ServiceContext
	dedup    *Dedup
	dedupTtl time.Duration
}

func New(sc *svc.ServiceContext) *Processor {
	return &Processor{
		sc:       sc,
		dedup:    NewDedup(sc.Redis),
		dedupTtl: time.Duration(sc.Config.Dedup.TtlSeconds) * time.Second,
	}
}

func (p *Processor) Handle(ctx context.Context, value []byte) error {
	flat, err := canalx.ParseFlat(value)
	if err != nil {
		logx.WithContext(ctx).Errorf("ParseFlat: %v", err)
		return nil
	}
	for _, row := range canalx.ExtractOutboxRows(flat) {
		if row.AggregateType != event.AggregateType {
			continue
		}
		if err := p.processRow(ctx, row); err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) processRow(ctx context.Context, row canalx.OutboxRow) error {
	dedupKey := fmt.Sprintf("dedup:rag:%s:%d", row.Type, row.Id)
	ok, err := p.dedup.Acquire(ctx, dedupKey, p.dedupTtl)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	var ev event.KnowPostEvent
	if err := json.Unmarshal([]byte(row.Payload), &ev); err != nil {
		logx.WithContext(ctx).Errorf("payload unmarshal: %v", err)
		return nil
	}
	switch ev.Type {
	case event.TypeKnowPostPublished, event.TypeKnowPostUpdated:
		return p.upsert(ctx, ev.PostId)
	case event.TypeKnowPostDeleted:
		return deleteAllChunks(ctx, p.sc.Es, p.sc.Config.RagIndex, ev.PostId)
	default:
		return nil
	}
}

// upsert 拉详情 + 拉正文 + 指纹比对 + 切块 + 嵌入 + 写 ES。
func (p *Processor) upsert(ctx context.Context, postId int64) error {
	d, err := p.sc.KnowPostRpc.GetDetail(ctx, &knowpostpb.GetDetailReq{Id: postId})
	if err != nil {
		return err
	}
	if d.Status != "published" || d.Visible != "public" {
		// 撤稿/转私 → 删除全部 chunk
		return deleteAllChunks(ctx, p.sc.Es, p.sc.Config.RagIndex, postId)
	}
	body, err := fetchContent(ctx, p.sc.HttpClient, d.ContentUrl)
	if err != nil || strings.TrimSpace(body) == "" {
		logx.WithContext(ctx).Errorf("rag fetch postId=%d err=%v len=%d", postId, err, len(body))
		// 拉正文失败：不阻塞 group；下一次事件再触发
		return nil
	}

	sha := fingerprint(body)
	ok, err := alreadyIndexed(ctx, p.sc.Es, p.sc.Config.RagIndex, postId, sha)
	if err != nil {
		return err
	}
	if ok {
		// 同 sha256 已索引过 → 跳过
		return nil
	}

	// 切块
	sections := textx.SplitByHeader(body)
	if len(sections) == 0 {
		sections = []textx.Section{{Body: body}}
	}
	type chunkMeta struct {
		Position int
		Title    string
		Text     string
	}
	var chunks []chunkMeta
	pos := 0
	for _, sec := range sections {
		for _, c := range textx.Chunk(sec.Body, p.sc.Config.Chunk.Size, p.sc.Config.Chunk.Overlap) {
			pos++
			chunks = append(chunks, chunkMeta{Position: pos, Title: sec.Title, Text: c})
		}
	}
	if len(chunks) == 0 {
		return nil
	}

	// 嵌入
	texts := make([]string, len(chunks))
	for i, c := range chunks {
		texts[i] = c.Text
	}
	// 嵌入（EINO DashScope，float64→float32 适配）
	vectors, err := llmx.EmbedFloat32(ctx, p.sc.Embed, texts)
	if err != nil {
		return err
	}
	if len(vectors) != len(chunks) {
		return fmt.Errorf("rag-indexer: vec count mismatch %d vs %d", len(vectors), len(chunks))
	}

	// 删旧版本（确保只有最新 sha 一份）
	if err := deleteAllChunks(ctx, p.sc.Es, p.sc.Config.RagIndex, postId); err != nil {
		return err
	}

	// bulk index
	var buf bytes.Buffer
	for i, c := range chunks {
		chunkId := fmt.Sprintf("p%d-%04d", postId, c.Position)
		actionLine, _ := json.Marshal(map[string]any{
			"index": map[string]any{
				"_index": p.sc.Config.RagIndex,
				"_id":    chunkId,
			},
		})
		doc := map[string]any{
			"post_id":        postId,
			"chunk_id":       chunkId,
			"position":       c.Position,
			"content_sha256": sha,
			"content_etag":   "", // 暂未拉 ETag；阶段4 简化
			"content_url":    d.ContentUrl,
			"title":          strings.TrimSpace(d.Title + " / " + c.Title),
			"text":           c.Text,
			"embedding":      vectors[i],
		}
		docLine, _ := json.Marshal(doc)
		buf.Write(actionLine)
		buf.WriteByte('\n')
		buf.Write(docLine)
		buf.WriteByte('\n')
	}
	if err := p.sc.Es.Bulk(ctx, buf.Bytes()); err != nil {
		return err
	}
	logx.WithContext(ctx).Infof("rag-indexer indexed postId=%d chunks=%d", postId, len(chunks))
	_ = strconv.Itoa(0) // keep strconv used (chunkId 已用 fmt)
	return nil
}

// fetchContent 简单 HTTP GET，返回正文。
func fetchContent(ctx context.Context, hc *http.Client, url string) (string, error) {
	if url == "" {
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
		return "", fmt.Errorf("fetch %s: %d", url, resp.StatusCode)
	}
	buf, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	return string(buf), err
}
