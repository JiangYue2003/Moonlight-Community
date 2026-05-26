package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/common/middleware"
	"github.com/zhiguang/zhiguang-go/services/llm/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 维持 llm 原生路由要求登录
	mw := middleware.NewAuthRpcMiddleware(serverCtx.AuthRpc, true)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{mw.Handle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "/describe", Handler: DescribeHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/qa/stream", Handler: QaStreamHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/api/v1/llm"),
	)

	// 前端兼容路由：可选鉴权（EventSource 默认不带 Authorization）
	optMw := middleware.NewAuthRpcMiddleware(serverCtx.AuthRpc, false)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{optMw.Handle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "/description/suggest", Handler: SuggestDescriptionCompatHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/:id/qa/stream", Handler: QaStreamCompatHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/api/v1/knowposts"),
	)
}
