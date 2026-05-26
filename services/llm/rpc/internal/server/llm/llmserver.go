package server

import (
	"context"

	llmlogic "github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/logic/llm"
	"github.com/zhiguang/zhiguang-go/services/llm/rpc/internal/svc"
	llmpb "github.com/zhiguang/zhiguang-go/services/llm/rpc/llm"
)

type LlmServer struct {
	svcCtx *svc.ServiceContext
	llmpb.UnimplementedLlmServer
}

func NewLlmServer(svcCtx *svc.ServiceContext) *LlmServer {
	return &LlmServer{svcCtx: svcCtx}
}

func (s *LlmServer) Describe(ctx context.Context, in *llmpb.DescribeReq) (*llmpb.DescribeResp, error) {
	return llmlogic.NewDescribeLogic(ctx, s.svcCtx).Describe(in)
}

func (s *LlmServer) QaStream(in *llmpb.QaStreamReq, stream llmpb.Llm_QaStreamServer) error {
	return llmlogic.NewQaStreamLogic(stream.Context(), s.svcCtx).Run(in, stream.Send)
}
