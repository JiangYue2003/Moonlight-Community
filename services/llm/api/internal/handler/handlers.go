package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/types"
)

// DescribeHandler 同步 POST 接口。
func DescribeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DescribeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewDescribeLogic(r.Context(), svcCtx)
		resp, err := l.Describe(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// QaStreamHandler SSE 自定义 writer，绕过 OkJson。
func QaStreamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewQaStreamLogic(r.Context(), svcCtx)
		// 任何错误已经在 logic 内部转为 SSE [DONE]，这里忽略
		_ = l.Run(w, &req)
	}
}

// SuggestDescriptionCompatHandler 前端兼容路径：
// POST /api/v1/knowposts/description/suggest  body: {content}
func SuggestDescriptionCompatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DescribeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewDescribeLogic(r.Context(), svcCtx)
		resp, err := l.Describe(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

// QaStreamCompatHandler 前端兼容路径：
// GET /api/v1/knowposts/{id}/qa/stream?question=...&topK=...&maxTokens=...
func QaStreamCompatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从兼容路径中提取 /api/v1/knowposts/{id}/qa/stream 的 {id}
		if req.PostId == 0 {
			path := strings.Trim(r.URL.Path, "/")
			parts := strings.Split(path, "/")
			// 期望: api v1 knowposts {id} qa stream
			if len(parts) >= 6 {
				if id, err := strconv.ParseInt(parts[3], 10, 64); err == nil {
					req.PostId = id
				}
			}
		}

		l := logic.NewQaStreamLogic(r.Context(), svcCtx)
		_ = l.Run(w, &req)
	}
}
