package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/pkg/llmx"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/prompt"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/types"
)
//
// 流程：
//  1. 验参；topK / maxTokens 边界
//  2. 取 question 的 embedding（EINO DashScope）
//  3. ES knnSearch 拿 fetchK 条；客户端按 postId 过滤 → topK
//  4. 0 命中 → 推 "未在该知文中找到相关内容。" + [DONE]
//  5. 拼上下文，EINO DeepSeek Stream，StreamReader.Recv() 逐 token 写 SSE
type QaStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQaStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QaStreamLogic {
	return &QaStreamLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

// Run 是 SSE 入口；写入由调用方 handler 完成 header 设置。
func (l *QaStreamLogic) Run(w http.ResponseWriter, req *types.QaReq) error {
	if req.PostId <= 0 || strings.TrimSpace(req.Question) == "" {
		return errorx.New(errorx.CodeBadRequest, "postId/question required")
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("streaming unsupported")
	}

	// per-user 限流（SSE 端点超限直接推错误并 [DONE]，不返回 HTTP 429）
	userId := userIdFromCtx(l.ctx)
	rlKey := fmt.Sprintf("rl:llm:qa:%d", userId)
	cfg := l.svcCtx.Config.LlmRateLimit
	allowed, rlErr := l.svcCtx.RateLimit.Take(l.ctx, rlKey, cfg.QaCapacity, cfg.QaRefillPerSec)
	if rlErr != nil {
		l.Logger.Errorf("ratelimit check error: %v", rlErr)
		allowed = true
	}
	if !allowed {
		w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.WriteHeader(http.StatusOK)
		flusher.Flush()
		l.send(w, flusher, "请求过于频繁，请稍后再试。")
		l.sendRaw(w, flusher, "[DONE]")
		return nil
	}

	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	topK := req.TopK
	if topK <= 0 || topK > 20 {
		topK = 5
	}
	maxTok := req.MaxTokens
	if maxTok <= 0 || maxTok > 4096 {
		maxTok = 1024
	}
	fetchK := topK * 3
	if fetchK < 20 {
		fetchK = 20
	}

	// 1. embedding（EINO DashScope，float64→float32 适配）
	vecs, err := llmx.EmbedFloat32(l.ctx, l.svcCtx.Embed, []string{req.Question})
	if err != nil || len(vecs) == 0 {
		return l.sseFatal(w, flusher, fmt.Errorf("embed failed: %v", err))
	}

	// 2. KNN 召回 + 过滤 postId
	res, err := l.svcCtx.Es.KnnSearch(l.ctx, l.svcCtx.Config.RagIndex, "embedding", vecs[0], fetchK, fetchK*2, nil)
	if err != nil {
		return l.sseFatal(w, flusher, err)
	}
	chunks := filterByPostId(res, req.PostId, topK)
	if len(chunks) == 0 {
		l.send(w, flusher, "未在该知文中找到相关内容。")
		l.sendRaw(w, flusher, "[DONE]")
		return nil
	}

	// 3. 拼 context
	var ctxBuf strings.Builder
	for i, c := range chunks {
		fmt.Fprintf(&ctxBuf, "[片段%d] %s\n", i+1, c)
	}
	user := fmt.Sprintf("基于以下知文上下文回答问题。\n\n上下文：\n%s\n\n问题：%s", ctxBuf.String(), strings.TrimSpace(req.Question))

	// 4. EINO Stream — StreamReader.Recv() + io.EOF 表示结束
	msgs := []*schema.Message{
		schema.SystemMessage(prompt.SysRag),
		schema.UserMessage(user),
	}
	reader, err := l.svcCtx.Chat.Stream(l.ctx, msgs)
	if err != nil {
		return l.sseFatal(w, flusher, err)
	}
	defer reader.Close()

	for {
		msg, err := reader.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return l.sseFatal(w, flusher, err)
		}
		if msg != nil && msg.Content != "" {
			l.send(w, flusher, msg.Content)
		}
	}
	l.sendRaw(w, flusher, "[DONE]")
	return nil
}

// send 推一条 data 行（自动转义换行）。
func (l *QaStreamLogic) send(w http.ResponseWriter, f http.Flusher, payload string) {
	safe := strings.ReplaceAll(payload, "\r", "")
	safe = strings.ReplaceAll(safe, "\n", "\\n")
	fmt.Fprintf(w, "data: %s\n\n", safe)
	f.Flush()
}

// sendRaw 不做转义（用于 [DONE] 这种特殊 token）。
func (l *QaStreamLogic) sendRaw(w http.ResponseWriter, f http.Flusher, payload string) {
	fmt.Fprintf(w, "data: %s\n\n", payload)
	f.Flush()
}

// sseFatal 向流推一条错误描述并 [DONE]，避免前端 EventSource 自动重连风暴。
func (l *QaStreamLogic) sseFatal(w http.ResponseWriter, f http.Flusher, err error) error {
	logx.WithContext(l.ctx).Errorf("qastream fatal: %v", err)
	l.send(w, f, "服务暂时不可用，请稍后再试。")
	l.sendRaw(w, f, "[DONE]")
	return nil
}

// filterByPostId 从 KNN 结果取 metadata.post_id == postId 的前 topK 条 text。
func filterByPostId(res *esx.SearchResult, postId int64, topK int) []string {
	if res == nil {
		return nil
	}
	type src struct {
		PostId int64  `json:"post_id"`
		Text   string `json:"text"`
	}
	out := make([]string, 0, topK)
	for _, h := range res.Hits.Hits {
		var s src
		if err := json.Unmarshal(h.Source, &s); err != nil {
			continue
		}
		if s.PostId != postId {
			continue
		}
		if strings.TrimSpace(s.Text) == "" {
			continue
		}
		out = append(out, s.Text)
		if len(out) >= topK {
			break
		}
	}
	return out
}
