package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	storageclient "github.com/zhiguang/zhiguang-go/services/storage/rpc/client/storage"
)

func registerMediaRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	storage := r.Group("/api/v1/storage", middleware.RequiredAuth(sc.AuthRpc))
	storage.POST("/presign", presign(sc))
}

func presign(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, ok := ctxdata.GetUserId(c.Request.Context())
		if !ok {
			gh.WriteError(c, errorx.New(errorx.CodeUnauthorized, "missing user id"))
			return
		}
		var req struct {
			Scene       string `json:"scene"`
			PostId      string `json:"postId"`
			ContentType string `json:"contentType"`
			Ext         string `json:"ext"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.StorageRpc.Presign(c.Request.Context(), &storageclient.PresignReq{
			UserId: uid, Scene: req.Scene, PostId: req.PostId, ContentType: req.ContentType, Ext: req.Ext,
		})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"url": resp.Url, "objectKey": resp.ObjectKey, "expiresIn": resp.ExpiresIn,
			"headers": resp.Headers, "contentUrl": resp.ContentUrl,
		})
	}
}
