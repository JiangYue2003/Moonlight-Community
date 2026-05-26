package llmlogic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/esx"
	"github.com/zhiguang/zhiguang-go/pkg/llmx"
	"github.com/zhiguang/zhiguang-go/services/llm/shared/prompt"
	"github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/svc"
	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
)

type QaStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQaStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QaStreamLogic {
	return &QaStreamLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *QaStreamLogic) Run(req *llmpb.QaStreamReq, send func(*llmpb.QaChunk) error) error {
	if req.PostId <= 0 || strings.TrimSpace(req.Question) == "" {
		return errorx.New(errorx.CodeBadRequest, "postId/question required")
	}

	rlKey := fmt.Sprintf("rl:llm:qa:%d", req.UserId)
	cfg := l.svcCtx.Config.LlmRateLimit
	allowed, rlErr := l.svcCtx.RateLimit.Take(l.ctx, rlKey, cfg.QaCapacity, cfg.QaRefillPerSec)
	if rlErr != nil {
		l.Logger.Errorf("ratelimit check error: %v", rlErr)
		allowed = true
	}
	if !allowed {
		if err := send(&llmpb.QaChunk{Data: "请求过于频繁，请稍后再试。"}); err != nil {
			return err
		}
		return send(&llmpb.QaChunk{Data: "[DONE]", Done: true})
	}

	topK := int(req.TopK)
	if topK <= 0 || topK > 20 {
		topK = 5
	}
	maxTok := int(req.MaxTokens)
	if maxTok <= 0 || maxTok > 4096 {
		maxTok = 1024
	}
	_ = maxTok
	fetchK := topK * 3
	if fetchK < 20 {
		fetchK = 20
	}

	vecs, err := llmx.EmbedFloat32(l.ctx, l.svcCtx.Embed, []string{req.Question})
	if err != nil || len(vecs) == 0 {
		return l.streamFatal(send, fmt.Errorf("embed failed: %v", err))
	}

	res, err := l.svcCtx.Es.KnnSearch(l.ctx, l.svcCtx.Config.RagIndex, "embedding", vecs[0], fetchK, fetchK*2, nil)
	if err != nil {
		return l.streamFatal(send, err)
	}
	chunks := filterByPostId(res, req.PostId, topK)
	if len(chunks) == 0 {
		if err := send(&llmpb.QaChunk{Data: "未在该知文中找到相关内容。"}); err != nil {
			return err
		}
		return send(&llmpb.QaChunk{Data: "[DONE]", Done: true})
	}

	var ctxBuf strings.Builder
	for i, c := range chunks {
		fmt.Fprintf(&ctxBuf, "[片段%d] %s\n", i+1, c)
	}
	user := fmt.Sprintf("基于以下知文上下文回答问题。\n\n上下文：\n%s\n\n问题：%s", ctxBuf.String(), strings.TrimSpace(req.Question))

	reader, err := l.svcCtx.Chat.Stream(l.ctx, []*schema.Message{
		schema.SystemMessage(prompt.SysRag),
		schema.UserMessage(user),
	})
	if err != nil {
		return l.streamFatal(send, err)
	}
	defer reader.Close()

	for {
		msg, err := reader.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return l.streamFatal(send, err)
		}
		if msg != nil && msg.Content != "" {
			if err := send(&llmpb.QaChunk{Data: sanitizeChunk(msg.Content)}); err != nil {
				return err
			}
		}
	}
	return send(&llmpb.QaChunk{Data: "[DONE]", Done: true})
}

func (l *QaStreamLogic) streamFatal(send func(*llmpb.QaChunk) error, err error) error {
	logx.WithContext(l.ctx).Errorf("qastream fatal: %v", err)
	if serr := send(&llmpb.QaChunk{Data: "服务暂时不可用，请稍后再试。"}); serr != nil {
		return serr
	}
	return send(&llmpb.QaChunk{Data: "[DONE]", Done: true})
}

func sanitizeChunk(payload string) string {
	safe := strings.ReplaceAll(payload, "\r", "")
	return strings.ReplaceAll(safe, "\n", "\\n")
}

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
		if s.PostId != postId || strings.TrimSpace(s.Text) == "" {
			continue
		}
		out = append(out, s.Text)
		if len(out) >= topK {
			break
		}
	}
	return out
}
