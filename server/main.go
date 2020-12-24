package main

import (
	"beau-blog/core"
	"beau-blog/global"
	"beau-blog/initialize"
)

func main() {
	// 初始化数据库
	global.BB_VP = core.Viper()          // 初始化viper
	global.BB_LOG = core.Zap()           // 日志
	global.BB_DB = initialize.Gorm() 		 // 连接数据库
	initialize.MysqlTables(global.BB_DB) // 表初始化, 自动迁移

	db, _ := global.BB_DB.DB()
	defer db.Close()

	core.RunWindowServer()
}
