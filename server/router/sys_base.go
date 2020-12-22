package router

func InitBaseRouter(Router *gin.RouterGroup) gin.IRoutes {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
