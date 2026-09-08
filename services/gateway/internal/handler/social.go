package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	usercounterclient "github.com/zhiguang/zhiguang-go/services/counter/rpc/client/usercounter"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	relationclient "github.com/zhiguang/zhiguang-go/services/relation/rpc/client/relation"
)

func registerSocialRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	relationPublic := r.Group("/api/v1/relation", middleware.OptionalAuth(sc.AuthRpc))
	relationPublic.GET("/status", relationStatus(sc))

	relationPrivate := r.Group("/api/v1/relation", middleware.RequiredAuth(sc.AuthRpc))
	{
		relationPrivate.POST("/follow", follow(sc))
		relationPrivate.POST("/unfollow", unfollow(sc))
		relationPrivate.GET("/following", listFollowing(sc))
		relationPrivate.GET("/followers", listFollowers(sc))
		relationPrivate.GET("/counter", relationCounter(sc))
	}
}

func follow(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		var req struct {
			ToUserId int64 `json:"toUserId"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err := sc.RelationRpc.Follow(c.Request.Context(), &relationclient.FollowReq{FromUserId: uid, ToUserId: req.ToUserId})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func unfollow(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		var req struct {
			ToUserId int64 `json:"toUserId"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		_, err := sc.RelationRpc.Unfollow(c.Request.Context(), &relationclient.UnfollowReq{FromUserId: uid, ToUserId: req.ToUserId})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func relationStatus(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		toUserId := int64(queryInt(c, "toUserId", 0))
		fromUserId, _ := ctxdata.GetUserId(c.Request.Context())
		resp, err := sc.RelationRpc.Status(c.Request.Context(), &relationclient.StatusReq{FromUserId: fromUserId, ToUserId: toUserId})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"following": resp.Following, "followedBy": resp.FollowedBy, "mutual": resp.Mutual})
	}
}

func listFollowing(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := int64(queryInt(c, "userId", 0))
		if userId == 0 {
			userId, _ = ctxdata.GetUserId(c.Request.Context())
		}
		limit := int32(queryInt(c, "limit", 20))
		offset := int32(queryInt(c, "offset", 0))
		cursor := int64(queryInt(c, "cursor", 0))
		resp, err := sc.RelationRpc.ListFollowing(c.Request.Context(), &relationclient.ListReq{UserId: userId, Limit: limit, Offset: offset, Cursor: cursor})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toRelationList(resp))
	}
}

func listFollowers(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := int64(queryInt(c, "userId", 0))
		if userId == 0 {
			userId, _ = ctxdata.GetUserId(c.Request.Context())
		}
		limit := int32(queryInt(c, "limit", 20))
		offset := int32(queryInt(c, "offset", 0))
		cursor := int64(queryInt(c, "cursor", 0))
		resp, err := sc.RelationRpc.ListFollowers(c.Request.Context(), &relationclient.ListReq{UserId: userId, Limit: limit, Offset: offset, Cursor: cursor})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, toRelationList(resp))
	}
}

func relationCounter(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := int64(queryInt(c, "userId", 0))
		if userId == 0 {
			userId, _ = ctxdata.GetUserId(c.Request.Context())
		}
		resp, err := sc.UserCounterRpc.GetUserSnapshot(c.Request.Context(), &usercounterclient.GetUserSnapshotReq{UserId: userId})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		snap := resp.GetSnapshot()
		c.JSON(http.StatusOK, gin.H{"followings": snap.GetFollowings(), "followers": snap.GetFollowers(), "posts": snap.GetPosts(), "likedPosts": int64(0), "favedPosts": int64(0), "likesReceived": snap.GetLikesReceived()})
	}
}

func toRelationList(resp *relationclient.ListResp) gin.H {
	items := make([]gin.H, 0, len(resp.Items))
	for _, item := range resp.Items {
		items = append(items, gin.H{"id": item.Id, "nickname": item.Nickname, "avatar": item.Avatar, "zgId": item.ZgId, "bio": item.Bio})
	}
	return gin.H{"items": items, "nextCursor": resp.NextCursor, "hasMore": resp.HasMore}
}
