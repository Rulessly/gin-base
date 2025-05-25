package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rulessly/gin-base/internal/router/system"
	"net/http"
)

type routerGroup struct {
	System system.RouterGroup
}

func Routers() *gin.Engine {
	group := new(routerGroup)
	router := gin.New()
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	//privateGroup := router.Group("/api")
	publicGroup := router.Group("/api")

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	systemRouter := group.System
	{
		systemRouter.InitBaseRouter(publicGroup)
	}

	return router
}
