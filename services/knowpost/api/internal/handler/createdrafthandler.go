// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

func CreateDraftHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCreateDraftLogic(r.Context(), svcCtx)
		resp, err := l.CreateDraft()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
