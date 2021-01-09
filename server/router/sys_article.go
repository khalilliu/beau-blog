package router

import (
	v1 "beau-blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAdminArticleRouter(Router *gin.RouterGroup) gin.IRoutes {
	articleRouter := Router.Group("article").Use()
	{
		articleRouter.GET("/list", v1.GetArticleList)
		articleRouter.GET("/:id", v1.GetArticleById)
		articleRouter.POST("/", v1.SaveArticle)
		articleRouter.PUT("/:id", v1.UpdateArticle)
		articleRouter.DELETE("/:id", v1.DeleteArticle)
	}
	return articleRouter
}


func InitClientArticleRouter(Router *gin.RouterGroup) gin.IRoutes {
	articleRouter := Router.GET("article")
	{
		articleRouter.GET("/list", v1.GetArticles)
		articleRouter.GET("/:id", v1.GetArticleById)
	}
	return articleRouter
}