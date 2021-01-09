package router

import (
	v1 "beau-blog/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAdminTagRouter(Router *gin.RouterGroup) gin.IRoutes {
		TagRouter := Router.Group("tag").Use()
		{
			TagRouter.GET("/:id", v1.GetTag)
			TagRouter.POST("/", v1.CreateTag)
			TagRouter.GET("edit/:id", v1.GetTag)
			TagRouter.PUT("/:id", v1.UpdateTag)
			TagRouter.DELETE("/:id", v1.DeleteTag)
		}
		return TagRouter
}