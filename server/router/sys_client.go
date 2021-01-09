package router

import (
	v1 "beau-blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitClientRouter(Router *gin.RouterGroup) gin.IRoutes {
	{
		// client 相关
		Router.GET("articles", v1.GetArticleList)
		Router.GET("article/:id", v1.GetArticleById)
	}
	return Router
}


