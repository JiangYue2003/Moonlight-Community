package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/common/middleware"
	"github.com/zhiguang/zhiguang-go/services/agent/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	mw := middleware.NewAuthRpcMiddleware(serverCtx.AuthRpc, true)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{mw.Handle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "/session", Handler: CreateSessionHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/chat/stream", Handler: ChatStreamHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/history", Handler: HistoryHandler(serverCtx)},
				{Method: http.MethodPost, Path: "/feedback", Handler: FeedbackHandler(serverCtx)},
				{Method: http.MethodPost, Path: "/memory/pin", Handler: MemoryPinHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/api/v1/agent"),
	)
}
