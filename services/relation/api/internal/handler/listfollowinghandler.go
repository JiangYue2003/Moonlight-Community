// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/relation/api/internal/types"
)

func ListFollowingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListFollowingLogic(r.Context(), svcCtx)
		resp, err := l.ListFollowing(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
