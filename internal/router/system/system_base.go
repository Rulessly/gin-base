package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/rulessly/gin-base/internal/api/v1"
)

type baseRouter struct{}

func (b *baseRouter) InitBaseRouter(routePublic *gin.RouterGroup) gin.IRoutes {

	baseRoutes := routePublic.Group("/system-base-v1")

	systemBaseApi := v1.ApiGroup.SystemApiGroup.BaseApi
	{
		baseRoutes.POST("login", systemBaseApi.Login)
		baseRoutes.POST("register", systemBaseApi.Register)
		baseRoutes.POST("logout", systemBaseApi.Logout)
		baseRoutes.GET("captcha", systemBaseApi.Captcha)
	}
	return baseRoutes
}
