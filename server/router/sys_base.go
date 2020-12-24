package router

import (
	v1 "beau-blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) gin.IRoutes {
	BaseRouter := Router.Group("base")
	{
		// admin 登录注册
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)

		// client 相关
		BaseRouter.GET("articles", v1.GetArticles)
		BaseRouter.GET("article/:id", v1.GetArticleById)
	}
	return BaseRouter
}
