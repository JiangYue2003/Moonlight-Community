package llm

import (
	"context"

	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DescribeReq  = llmpb.DescribeReq
	DescribeResp = llmpb.DescribeResp
	QaStreamReq  = llmpb.QaStreamReq
	QaChunk      = llmpb.QaChunk

	Llm interface {
		Describe(ctx context.Context, in *DescribeReq, opts ...grpc.CallOption) (*DescribeResp, error)
		QaStream(ctx context.Context, in *QaStreamReq, opts ...grpc.CallOption) (llmpb.Llm_QaStreamClient, error)
	}

	defaultLlm struct {
		cli zrpc.Client
	}
)

func NewLlm(cli zrpc.Client) Llm {
	return &defaultLlm{cli: cli}
}

func (m *defaultLlm) Describe(ctx context.Context, in *DescribeReq, opts ...grpc.CallOption) (*DescribeResp, error) {
	client := llmpb.NewLlmClient(m.cli.Conn())
	return client.Describe(ctx, in, opts...)
}

func (m *defaultLlm) QaStream(ctx context.Context, in *QaStreamReq, opts ...grpc.CallOption) (llmpb.Llm_QaStreamClient, error) {
	client := llmpb.NewLlmClient(m.cli.Conn())
	return client.QaStream(ctx, in, opts...)
}
