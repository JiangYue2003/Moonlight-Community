package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/eino/schema"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/pkg/textx"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/prompt"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/types"
)

type DescribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDescribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DescribeLogic {
	return &DescribeLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

// Describe 生成 ≤50 个汉字的描述（NFKC + 折空白 + 去引号 + 去末标点 + 截断）。
func (l *DescribeLogic) Describe(req *types.DescribeReq) (*types.DescribeResp, error) {
	body := strings.TrimSpace(req.Body)
	if body == "" {
		body = strings.TrimSpace(req.Content)
	}
	if body == "" {
		return nil, errorx.New(errorx.CodeBadRequest, "body required")
	}

	// per-user 限流（userId 由 JWT 中间件注入到 ctx）
	userId := userIdFromCtx(l.ctx)
	rlKey := fmt.Sprintf("rl:llm:describe:%d", userId)
	cfg := l.svcCtx.Config.LlmRateLimit
	allowed, err := l.svcCtx.RateLimit.Take(l.ctx, rlKey, cfg.DescribeCapacity, cfg.DescribeRefillPerSec)
	if err != nil {
		l.Logger.Errorf("ratelimit check error: %v", err)
		allowed = true // Redis 故障时放行，避免限流组件拖垮功能
	}
	if !allowed {
		return nil, errorx.New(errorx.CodeRateLimited, "请求过于频繁，请稍后再试")
	}

	msgs := []*schema.Message{
		schema.SystemMessage(prompt.SysDescribe),
		schema.UserMessage(textx.TruncateRunes(body, 4000)),
	}
	resp, err := l.svcCtx.Chat.Generate(l.ctx, msgs)
	if err != nil {
		return nil, err
	}

	return &types.DescribeResp{
		Description: textx.DescriptionPostProcess{MaxRunes: 50}.Apply(resp.Content),
	}, nil
}

// userIdFromCtx 从 go-zero JWT 中间件注入的 context 取 userId；未登录返回 0。
func userIdFromCtx(ctx context.Context) int64 {
	uid, _ := ctxdata.GetUserId(ctx)
	return uid
}
