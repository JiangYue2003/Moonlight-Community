package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	searchclient "github.com/zhiguang/zhiguang-go/services/search/rpc/client/search"
)

func registerDiscoveryRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	search := r.Group("/api/v1/search", middleware.OptionalAuth(sc.AuthRpc))
	{
		search.GET("/", searchPosts(sc))
		search.GET("/suggest", suggestSearch(sc))
	}
}

func searchPosts(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		viewer, _ := ctxdata.GetUserId(c.Request.Context())
		size := int32(queryInt(c, "size", 20))
		resp, err := sc.SearchRpc.Search(c.Request.Context(), &searchclient.SearchReq{Q: c.Query("q"), Size: size, Tags: c.Query("tags"), After: c.Query("after"), ViewerId: viewer})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		items := make([]gin.H, 0, len(resp.Items))
		for _, item := range resp.Items {
			items = append(items, gin.H{"contentId": item.Id, "contentType": "knowpost", "title": item.Title, "description": item.Description, "snippet": item.Description, "tags": item.Tags, "authorId": item.AuthorId, "authorNickname": item.AuthorNickname, "authorAvatar": item.AuthorAvatar, "likeCount": item.LikeCount, "favoriteCount": item.FavoriteCount, "viewCount": 0, "imgUrls": []string{item.CoverImage}, "isTop": item.IsTop})
		}
		c.JSON(http.StatusOK, gin.H{"items": items, "nextAfter": resp.NextAfter, "hasMore": resp.HasMore})
	}
}

func suggestSearch(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		size := int32(queryInt(c, "size", 10))
		resp, err := sc.SearchRpc.Suggest(c.Request.Context(), &searchclient.SuggestReq{Prefix: c.Query("prefix"), Size: size})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"items": resp.Items})
	}
}
