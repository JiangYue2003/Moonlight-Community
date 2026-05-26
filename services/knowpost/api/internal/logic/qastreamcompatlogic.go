package logic

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

type QaStreamCompatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	postID int64
}

func NewQaStreamCompatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QaStreamCompatLogic {
	return &QaStreamCompatLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *QaStreamCompatLogic) WithPostID(id int64) *QaStreamCompatLogic {
	l.postID = id
	return l
}

func (l *QaStreamCompatLogic) Run(w http.ResponseWriter, r *http.Request) error {
	if l.postID <= 0 {
		return errorx.New(errorx.CodeBadRequest, "invalid post id")
	}
	base := strings.TrimRight(l.svcCtx.Config.LlmProxy.BaseURL, "/")
	if base == "" {
		return errorx.New(errorx.CodeInternalError, "llm proxy base url required")
	}

	q := r.URL.Query()
	q.Set("postId", strconv.FormatInt(l.postID, 10))
	u, err := url.Parse(base + "/api/v1/llm/qa/stream")
	if err != nil {
		return err
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(l.ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
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
