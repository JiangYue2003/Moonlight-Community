package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	counterclient "github.com/zhiguang/zhiguang-go/services/counter/rpc/client/counter"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
)

func registerEngagementRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	action := r.Group("/api/v1/action", middleware.RequiredAuth(sc.AuthRpc))
	{
		action.POST("/like", toggleMetric(sc, "like", true))
		action.POST("/unlike", toggleMetric(sc, "like", false))
		action.POST("/fav", toggleMetric(sc, "fav", true))
		action.POST("/unfav", toggleMetric(sc, "fav", false))
	}
	r.GET("/api/v1/counter/:etype/:eid", getCounts(sc))
}

func toggleMetric(sc *srv.ServiceContext, metric string, add bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		var req struct {
			EntityType string `json:"entityType"`
			EntityId   string `json:"entityId"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.CounterRpc.Toggle(c.Request.Context(), &counterclient.ToggleReq{EntityType: req.EntityType, EntityId: req.EntityId, Metric: metric, UserId: uid, Add: add})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		key := map[string]string{"like": "liked", "fav": "faved"}[metric]
		c.JSON(http.StatusOK, gin.H{"changed": resp.Changed, key: add})
	}
}

func getCounts(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		etype := c.Param("etype")
		eid := c.Param("eid")
		resp, err := sc.CounterRpc.GetCounts(c.Request.Context(), &counterclient.GetCountsReq{EntityType: etype, EntityId: eid})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"entityType": etype, "entityId": eid, "counts": resp.Counts})
	}
}
