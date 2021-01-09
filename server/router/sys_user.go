package router

import (
	v1 "beau-blog/api/v1"
	"beau-blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdminUserRouter(Router *gin.RouterGroup) gin.IRoutes {
	UserRouter := Router.Group("user").Use(middleware.CurrentUser())
	{
		UserRouter.POST("changePassword", v1.ChangePassword)
		UserRouter.PUT("setUserInfo", v1.SetUserInfo)
	}
	return UserRouter
}
