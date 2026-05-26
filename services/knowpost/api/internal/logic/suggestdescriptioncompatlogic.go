package logic

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

type SuggestDescriptionCompatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuggestDescriptionCompatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuggestDescriptionCompatLogic {
	return &SuggestDescriptionCompatLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *SuggestDescriptionCompatLogic) Run(w http.ResponseWriter, r *http.Request) error {
	base := strings.TrimRight(l.svcCtx.Config.LlmProxy.BaseURL, "/")
	if base == "" {
		return errorx.New(errorx.CodeInternalError, "llm proxy base url required")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	_ = r.Body.Close()

	req, err := http.NewRequestWithContext(l.ctx, http.MethodPost, base+"/api/v1/llm/describe", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if auth := r.Header.Get("Authorization"); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	resp, err := l.svcCtx.LlmClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	return err
}
