package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
)

func parsePathInt64(c *gin.Context, name string) (int64, error) {
	v, err := strconv.ParseInt(c.Param(name), 10, 64)
	if err != nil {
		return 0, errorx.New(errorx.CodeBadRequest, "invalid "+name)
	}
	return v, nil
}

func queryInt(c *gin.Context, key string, def int) int {
	if s := c.Query(key); s != "" {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
	}
	return def
}

func queryIntStrict(c *gin.Context, key string, def int) (int, error) {
	s := c.Query(key)
	if s == "" {
		return def, nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, errorx.New(errorx.CodeBadRequest, "invalid "+key)
	}
	return n, nil
}
