// 由人工修改：将 rest.WithJwt 替换为自定义 AuthRpcMiddleware（RS256 走 auth-rpc.VerifyToken）。

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zhiguang/zhiguang-go/common/middleware"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	authMw := middleware.NewAuthRpcMiddleware(serverCtx.AuthRpc, true)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMw.Handle},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/fav",
					Handler: FavHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/like",
					Handler: LikeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/unfav",
					Handler: UnfavHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/unlike",
					Handler: UnlikeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/action"),
	)

	// /counter/:etype/:eid 受保护接口，按文档约定要求鉴权。
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{authMw.Handle},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:etype/:eid",
					Handler: GetCountsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/counter"),
	)
}
