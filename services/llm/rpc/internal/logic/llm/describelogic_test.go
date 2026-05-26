package llmlogic

import (
	"context"
	"strings"
	"testing"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/cloudwego/eino/components/model"
	goredis "github.com/redis/go-redis/v9"
	"github.com/cloudwego/eino/schema"
	"github.com/zhiguang/zhiguang-go/pkg/ratelimit"
	"github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/config"
	"github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/svc"
	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
)

type stubChatModel struct{}

func (stubChatModel) Generate(ctx context.Context, msgs []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	return &schema.Message{Content: "  这是一个自动生成的摘要。  "}, nil
}

func (stubChatModel) Stream(ctx context.Context, msgs []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	return nil, nil
}

func (stubChatModel) BindTools(tools []*schema.ToolInfo) error {
	return nil
}

func newTestRateLimit(t *testing.T) *ratelimit.TokenBucket {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mr.Close)
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{Addrs: []string{mr.Addr()}})
	t.Cleanup(func() { _ = rdb.Close() })
	return ratelimit.New(rdb)
}

func TestDescribe_RejectsEmptyBody(t *testing.T) {
	sc := &svc.ServiceContext{Config: config.Config{}, RateLimit: newTestRateLimit(t)}
	_, err := NewDescribeLogic(context.Background(), sc).Describe(&llmpb.DescribeReq{})
	if err == nil {
		t.Fatal("empty body should fail")
	}
}

func TestDescribe_UsesContentFallback(t *testing.T) {
	sc := &svc.ServiceContext{
		Config: config.Config{
			LlmRateLimit: config.LlmRateLimitConf{
				DescribeCapacity:     100,
				DescribeRefillPerSec: 100,
			},
		},
		Chat:      stubChatModel{},
		RateLimit: newTestRateLimit(t),
	}
	resp, err := NewDescribeLogic(context.Background(), sc).Describe(&llmpb.DescribeReq{Content: strings.Repeat("x", 20)})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(resp.Description, "自动生成") {
		t.Fatalf("unexpected description=%q", resp.Description)
	}
}
