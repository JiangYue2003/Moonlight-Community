package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
)

// NewEngine is the Gateway HTTP composition root. Bounded contexts own their
// route registration and handlers in sibling files in this package.
func NewEngine(sc *srv.ServiceContext) *gin.Engine {
	r := gin.New()
	if sc.Config.HTTPAccessLog {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())

	registerIdentityRoutes(r, sc)
	registerMediaRoutes(r, sc)
	registerContentRoutes(r, sc)
	registerSocialRoutes(r, sc)
	registerEngagementRoutes(r, sc)
	registerDiscoveryRoutes(r, sc)
	registerIntelligenceRoutes(r, sc)

	return r
}
