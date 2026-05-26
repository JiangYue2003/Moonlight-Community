package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

func GetDetailPublicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := extractIdFromPath(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		resp, err := logic.NewGetDetailPublicLogic(r.Context(), svcCtx).WithId(id).GetDetailPublic()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
