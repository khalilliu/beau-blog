package router

func InitUserRouter(Router *gin.RouterGroup) gin.IRoutes {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)
		UserRouter.POST("changePassword", v1.ChangePassword)
		UserRouter.DELETE("deleteUser", v1.DeleteUser)
		UserRouter.PUT("setUserInfo", v1.SetUserInfo)
	}
	return UserRouter
}
