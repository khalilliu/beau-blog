package initialize

import (
	"beau-blog/global"
	"beau-blog/middleware"
	"beau-blog/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.BB_CONFIG.Local.Path, http.Dir(global.BB_CONFIG.Local.Path)) // 静态文件
	global.BB_LOG.Info("use middleware logger")

	// 跨域
	Router.Use(middleware.Cors())
	global.BB_LOG.Info("use middleware cors")
	global.BB_LOG.Info("register swagger handler")
	// 客户端api
	clientGroup := Router.Group("client")
	{
		router.InitClientRouter(clientGroup)
	}
	// 服务端api
	adminGroup := Router.Group("admin")
	{
		router.InitAdminRouter(adminGroup)
	}
	global.BB_LOG.Info("router register success")
	return Router
}
