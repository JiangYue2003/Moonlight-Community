package llmlogic

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/textx"
	"github.com/zhiguang/zhiguang-go/services/llm/shared/prompt"
	"github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/svc"
	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
)

type DescribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDescribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DescribeLogic {
	return &DescribeLogic{ctx: ctx, svcCtx: svcCtx, Logger: logx.WithContext(ctx)}
}

func (l *DescribeLogic) Describe(req *llmpb.DescribeReq) (*llmpb.DescribeResp, error) {
	body := strings.TrimSpace(req.Body)
	if body == "" {
		body = strings.TrimSpace(req.Content)
	}
	if body == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "body required")
	}

	rlKey := fmt.Sprintf("rl:llm:describe:%d", req.UserId)
	cfg := l.svcCtx.Config.LlmRateLimit
	allowed, err := l.svcCtx.RateLimit.Take(l.ctx, rlKey, cfg.DescribeCapacity, cfg.DescribeRefillPerSec)
	if err != nil {
		l.Logger.Errorf("ratelimit check error: %v", err)
		allowed = true
	}
	if !allowed {
		return nil, errorx.New(errorx.CodeRateLimited, "请求过于频繁，请稍后再试")
	}

	resp, err := l.svcCtx.Chat.Generate(l.ctx, []*schema.Message{
		schema.SystemMessage(prompt.SysDescribe),
		schema.UserMessage(textx.TruncateRunes(body, 4000)),
	})
	if err != nil {
		return nil, err
	}
	return &llmpb.DescribeResp{
		Description: textx.DescriptionPostProcess{MaxRunes: 50}.Apply(resp.Content),
	}, nil
}
