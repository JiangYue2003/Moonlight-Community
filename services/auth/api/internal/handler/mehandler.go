// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/logic"
	"github.com/zhiguang/zhiguang-go/services/auth/api/internal/svc"
)

func MeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := ""
		const prefix = "Bearer "
		if h := r.Header.Get("Authorization"); strings.HasPrefix(h, prefix) {
			token = strings.TrimSpace(h[len(prefix):])
		}
		l := logic.NewMeLogic(r.Context(), svcCtx).WithToken(token)
		resp, err := l.Me()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
