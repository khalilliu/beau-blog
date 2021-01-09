package router

import "github.com/gin-gonic/gin"

func InitAdminCateRouter(Router *gin.RouterGroup) gin.IRoutes {
	CateRouter := Router.Group("category")
	{
		CateRouter.GET("/:id", v1.GetCategory)
		CateRouter.POST("/", v1.CreateCategory)
		CateRouter.GET("edit/:id", v1.GetCategory)
		CateRouter.PUT("/:id", v1.UpdateCategory)
		CateRouter.DELETE("/:id", v1.DeleteCategory)
	}
	return CateRouter
}