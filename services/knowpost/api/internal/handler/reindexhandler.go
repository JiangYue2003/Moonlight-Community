package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

func ReindexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := extractIdFromPath(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		resp, err := logic.NewReindexLogic(r.Context(), svcCtx).WithId(id).Reindex()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			_ = resp
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
