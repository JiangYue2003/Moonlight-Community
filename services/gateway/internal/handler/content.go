package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	knowpostclient "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/client/knowpost"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxFeedCursorLength = 256

func registerContentRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	publicKnowpost := r.Group("/api/v1/knowposts", middleware.OptionalAuth(sc.AuthRpc))
	{
		publicKnowpost.GET("/feed", getPublicFeed(sc))
		publicKnowpost.GET("/detail/:id", getDetail(sc))
		publicKnowpost.POST("/description/suggest", suggestDescription(sc))
		publicKnowpost.GET("/:id/qa/stream", qaCompatStream(sc))
	}

	privateKnowpost := r.Group("/api/v1/knowposts", middleware.RequiredAuth(sc.AuthRpc))
	{
		privateKnowpost.POST("/drafts", createDraft(sc))
		privateKnowpost.POST("/:id/content/confirm", confirmContent(sc))
		privateKnowpost.PATCH("/:id", patchMetadata(sc))
		privateKnowpost.POST("/:id/publish", publish(sc))
		privateKnowpost.PATCH("/:id/top", updateTop(sc))
		privateKnowpost.PATCH("/:id/visibility", updateVisibility(sc))
		privateKnowpost.DELETE("/:id", deleteKnowpost(sc))
		privateKnowpost.GET("/mine", getMyFeed(sc))
		privateKnowpost.GET("/following-feed", getFollowingFeed(sc))
		privateKnowpost.POST("/:id/reindex", reindex(sc))
		privateKnowpost.POST("/:id/rag/reindex", reindex(sc))
	}
}

func createDraft(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		resp, err := sc.KnowPostRpc.CreateDraft(c.Request.Context(), &knowpostclient.CreateDraftReq{CreatorId: uid})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": resp.Id})
	}
}

func confirmContent(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		var req struct {
			ObjectKey string `json:"objectKey"`
			Etag      string `json:"etag"`
			Size      int64  `json:"size"`
			Sha256    string `json:"sha256"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err = sc.KnowPostRpc.ConfirmContent(c.Request.Context(), &knowpostclient.ConfirmContentReq{Id: id, CreatorId: uid, ObjectKey: req.ObjectKey, Etag: req.Etag, Size: req.Size, Sha256: req.Sha256})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func patchMetadata(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		var req struct {
			Title       *string  `json:"title"`
			Description *string  `json:"description"`
			TagId       *int64   `json:"tagId"`
			Tags        []string `json:"tags"`
			TagsSet     bool     `json:"tagsSet"`
			ImgUrls     []string `json:"imgUrls"`
			ImgUrlsSet  bool     `json:"imgUrlsSet"`
			Visible     *string  `json:"visible"`
			IsTop       *bool    `json:"isTop"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		in := &knowpostclient.PatchMetadataReq{Id: id, CreatorId: uid}
		if req.Title != nil {
			in.Title, in.TitleSet = *req.Title, true
		}
		if req.Description != nil {
			in.Description, in.DescriptionSet = *req.Description, true
		}
		if req.TagId != nil {
			in.TagId, in.TagIdSet = *req.TagId, true
		}
		in.Tags, in.TagsSet = req.Tags, req.TagsSet
		in.ImgUrls, in.ImgUrlsSet = req.ImgUrls, req.ImgUrlsSet
		if req.Visible != nil {
			in.Visible, in.VisibleSet = *req.Visible, true
		}
		if req.IsTop != nil {
			in.IsTop, in.IsTopSet = *req.IsTop, true
		}
		resp, err := sc.KnowPostRpc.PatchMetadata(c.Request.Context(), in)
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toKnowPostDetail(resp))
	}
}

func publish(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		resp, err := sc.KnowPostRpc.Publish(c.Request.Context(), &knowpostclient.PublishReq{Id: id, CreatorId: uid})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toKnowPostDetail(resp))
	}
}

func updateTop(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		var req struct {
			IsTop bool `json:"isTop"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err = sc.KnowPostRpc.UpdateTop(c.Request.Context(), &knowpostclient.UpdateTopReq{Id: id, CreatorId: uid, IsTop: req.IsTop})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func updateVisibility(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		var req struct {
			Visible string `json:"visible"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err = sc.KnowPostRpc.UpdateVisibility(c.Request.Context(), &knowpostclient.UpdateVisibilityReq{Id: id, CreatorId: uid, Visible: req.Visible})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func deleteKnowpost(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		_, err = sc.KnowPostRpc.Delete(c.Request.Context(), &knowpostclient.DeleteReq{Id: id, CreatorId: uid})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func getPublicFeed(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := int32(queryInt(c, "page", 1))
		size := int32(queryInt(c, "size", 20))
		resp, err := sc.KnowPostRpc.GetPublicFeed(c.Request.Context(), &knowpostclient.GetPublicFeedReq{Page: page, Size: size})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toFeedPage(resp))
	}
}

func getMyFeed(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		page := int32(queryInt(c, "page", 1))
		size := int32(queryInt(c, "size", 20))
		resp, err := sc.KnowPostRpc.GetMyFeed(c.Request.Context(), &knowpostclient.GetMyFeedReq{CreatorId: uid, Page: page, Size: size})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toFeedPage(resp))
	}
}

func getFollowingFeed(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		cursor := c.Query("cursor")
		if len(cursor) > maxFeedCursorLength {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "cursor is too long"))
			return
		}
		sizeValue, err := queryIntStrict(c, "size", 20)
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		if sizeValue < 1 || sizeValue > 100 {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "size must be between 1 and 100"))
			return
		}
		page := int32(0)
		if cursor == "" {
			pageValue, err := queryIntStrict(c, "page", 1)
			if err != nil {
				gh.WriteError(c, err)
				return
			}
			if pageValue < 1 {
				gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "page must be >= 1"))
				return
			}
			page = int32(pageValue)
		}
		size := int32(sizeValue)
		resp, err := sc.KnowPostRpc.GetUserFeed(c.Request.Context(), &knowpostclient.GetUserFeedReq{
			UserId: uid,
			Page:   page,
			Size:   size,
			Cursor: cursor,
		})
		if err != nil {
			if status.Code(err) == codes.InvalidArgument {
				gh.WriteError(c, errorx.New(errorx.CodeBadRequest, status.Convert(err).Message()))
				return
			}
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toFeedPage(resp))
	}
}

func getDetail(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		viewer, _ := ctxdata.GetUserId(c.Request.Context())
		resp, err := sc.KnowPostRpc.GetDetail(c.Request.Context(), &knowpostclient.GetDetailReq{Id: id, ViewerId: viewer})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toKnowPostDetail(resp))
	}
}

func reindex(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		id, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		_, err = sc.KnowPostRpc.Reindex(c.Request.Context(), &knowpostclient.ReindexReq{Id: id, CreatorId: uid})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func toKnowPostDetail(resp *knowpostclient.KnowPostDetail) gin.H {
	return gin.H{
		"id": resp.Id, "creatorId": resp.CreatorId, "title": resp.Title, "description": resp.Description, "tagId": resp.TagId,
		"tags": resp.Tags, "contentUrl": resp.ContentUrl, "contentObjectKey": resp.ContentObjectKey, "imgUrls": resp.ImgUrls,
		"visible": resp.Visible, "status": resp.Status, "type": resp.Type, "isTop": resp.IsTop,
		"createTime": resp.CreateTime, "updateTime": resp.UpdateTime, "publishTime": resp.PublishTime,
	}
}

func toFeedPage(resp *knowpostclient.FeedPage) gin.H {
	items := make([]gin.H, 0, len(resp.Items))
	for _, item := range resp.Items {
		items = append(items, gin.H{
			"id": item.Id, "creatorId": item.CreatorId, "title": item.Title, "description": item.Description,
			"contentUrl": item.ContentUrl, "tags": item.Tags, "imgUrls": item.ImgUrls, "visible": item.Visible,
			"isTop": item.IsTop, "publishTime": item.PublishTime,
		})
	}
	return gin.H{
		"items": items, "hasMore": resp.HasMore, "size": resp.Size, "page": resp.Page,
		"nextCursor": resp.NextCursor,
	}
}
