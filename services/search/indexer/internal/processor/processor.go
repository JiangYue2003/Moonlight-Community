// Package processor search-indexer 主处理器：从 canal-outbox 抽 KnowPostEvent，更新 ES 索引。
package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/canalx"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	event "github.com/zhiguang/zhiguang-go/services/knowpost/shared/event"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/svc"
)

// Processor 顶层入口。
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

// Handle 处理一条 Kafka 消息。
func (p *Processor) Handle(ctx context.Context, value []byte) error {
	flat, err := canalx.ParseFlat(value)
	if err != nil {
		logx.WithContext(ctx).Errorf("canalx ParseFlat: %v", err)
		return nil
	}
	rows := canalx.ExtractOutboxRows(flat)
	for _, row := range rows {
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
	dedupKey := fmt.Sprintf("dedup:idx:%s:%d", row.Type, row.Id)
	ok, err := p.dedup.Acquire(ctx, dedupKey, p.dedupTtl)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	var ev event.KnowPostEvent
	if err := json.Unmarshal([]byte(row.Payload), &ev); err != nil {
		logx.WithContext(ctx).Errorf("payload unmarshal: %v body=%s", err, row.Payload)
		return nil
	}

	switch ev.Type {
	case event.TypeKnowPostPublished, event.TypeKnowPostUpdated:
		return p.upsert(ctx, ev.PostId)
	case event.TypeKnowPostDeleted:
		return p.softDelete(ctx, ev.PostId)
	default:
		// KnowPostCreated 等草稿态事件不进入索引
		return nil
	}
}

// upsert 拉详情 + 拉正文 + 写 ES。
//
// 若 detail 显示已不再 published/public（用户改私 / 撤稿），转为 SoftDelete。
func (p *Processor) upsert(ctx context.Context, postId int64) error {
	detail, err := p.sc.KnowPostRpc.GetDetail(ctx, &knowpostpb.GetDetailReq{Id: postId})
	if err != nil {
		return err
	}
	if detail.Status != "published" || detail.Visible != "public" {
		return p.softDelete(ctx, postId)
	}
	body, err := fetchContent(ctx, p.sc.HttpClient, detail.ContentUrl, p.sc.Config.ContentMaxRunes)
	if err != nil {
		// 拉正文失败：跳过本条但记录；不阻塞 group 进度
		logx.WithContext(ctx).Errorf("fetch contentUrl failed postId=%d: %v", postId, err)
		body = ""
	}
	doc := buildDoc(detail, body)
	return p.sc.Es.Index(ctx, p.sc.Config.ContentIndex, strconv.FormatInt(postId, 10), doc)
}

func (p *Processor) softDelete(ctx context.Context, postId int64) error {
	partial := map[string]any{"doc": map[string]any{"status": "deleted"}}
	if err := p.sc.Es.Update(ctx, p.sc.Config.ContentIndex, strconv.FormatInt(postId, 10), partial); err != nil {
		// 文档不存在不视为错误
		logx.WithContext(ctx).Infof("softDelete update postId=%d: %v", postId, err)
	}
	return nil
}

// buildDoc 把 KnowPostDetail + 正文转换为搜索文档。
func buildDoc(d *knowpostpb.KnowPostDetail, body string) map[string]any {
	pubMs := d.PublishTime
	return map[string]any{
		"content_id":     d.Id,
		"content_type":   d.Type,
		"title":          strings.TrimSpace(d.Title),
		"body":           body,
		"description":    d.Description,
		"tags":           d.Tags,
		"img_urls":       d.ImgUrls,
		"author_id":      d.CreatorId,
		"publish_time":   pubMs,
		"status":         d.Status,
		"is_top":         d.IsTop,
		"like_count":     0, // counter 走另条路径维护，此处先置 0
		"favorite_count": 0,
		"view_count":     0,
		"title_suggest": []map[string]any{
			{"input": d.Title},
		},
	}
}
