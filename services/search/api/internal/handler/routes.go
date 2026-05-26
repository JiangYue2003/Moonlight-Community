package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/zhiguang/zhiguang-go/common/middleware"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/svc"
)

// RegisterHandlers 注册路由。
//
// 搜索接口的 JWT 是**可选**的：required=false 表示无 token 时仍能匿名访问。
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	mw := middleware.NewAuthRpcMiddleware(serverCtx.AuthRpc, false)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{mw.Handle},
			[]rest.Route{
				{Method: http.MethodGet, Path: "/", Handler: SearchHandler(serverCtx)},
				{Method: http.MethodGet, Path: "/suggest", Handler: SuggestHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/api/v1/search"),
	)
}
