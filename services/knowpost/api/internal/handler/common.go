// Package handler 共享工具：提取 :id path 参数。
package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/pathvar"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
)

func extractIdFromPath(r *http.Request) (int64, error) {
	vars := pathvar.Vars(r)
	s, ok := vars["id"]
	if !ok || s == "" {
		return 0, errorx.New(errorx.CodeBadRequest, "missing :id in path")
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errorx.Wrap(errorx.CodeBadRequest, ":id not numeric", err)
	}
	return v, nil
}
