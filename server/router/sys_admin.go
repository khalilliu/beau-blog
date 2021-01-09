package router

import (
	v1 "beau-blog/api/v1"
	"beau-blog/middleware"
	"github.com/gin-gonic/gin"
)


func InitAdminRouter(Router *gin.RouterGroup) gin.IRoutes {
	{
		// 不需要登录
		Router.GET("login", v1.Login)
		Router.GET("captcha", v1.Captcha)
	}

	Router.Use(middleware.JWTAuth(), middleware.CurrentUser())
	{
		/**
		* 用户相关路由
		* /admin/user/
		*/
		InitAdminUserRouter(Router)
		InitAdminArticleRouter(Router)
		InitAdminTagRouter(Router)
		InitAdminCateRouter(Router)
	}
	return Router
}
