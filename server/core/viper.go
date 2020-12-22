package core

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Server interface {
}

func RunWindowServer() {
	if global.BB_CONFIG.System.UseMultipoint {
		// 初始化redis
		initialize.Redis()
	}
	initialize.InitWkMode()
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.BB_CONFIG.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.BB_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
		默认前端文件运行地址:http://127.0.0.1:8082
	`)

	global.BB_LOG.Error(s.ListenAndServe().Error())
}
