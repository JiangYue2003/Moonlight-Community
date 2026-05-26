package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
)

func SuggestDescriptionCompatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSuggestDescriptionCompatLogic(r.Context(), svcCtx)
		if err := l.Run(w, r); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}

func QaStreamCompatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil || id <= 0 {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewQaStreamCompatLogic(r.Context(), svcCtx).WithPostID(id)
		if err := l.Run(w, r); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}
