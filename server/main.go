package main

func main() {
	// 初始化数据库
	global.BB_VP = core.Viper()          // 初始化viper
	global.BB_LOG = core.Zap()           //日志
	initialize.MysqlTables(global.BB_DB) // 数据库

	db, _ := global.BB_DB.DB()
	defer db.Close()

	core.RunWindowServer()
}
