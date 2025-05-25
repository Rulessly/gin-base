package system

import "github.com/gin-gonic/gin"

type baseRouter struct{}

func (b *baseRouter) InitBaseRouter(RouterGroup *gin.RouterGroup) gin.IRoutes {
	baseRoutes := RouterGroup.Group("base")
	{
		baseRoutes.POST("login", nil)
		baseRoutes.POST("captcha", nil)
	}
	return baseRoutes
}
