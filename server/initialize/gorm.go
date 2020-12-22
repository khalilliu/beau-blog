package initialize

import (
	"beau-blog/model"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@author: Khalil
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func Gorm() *gorm.DB {
	switch global.BB_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: Khalil
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
	)
	if err != nil {
		global.BB_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.BB_LOG.Info("register table success")
}

//@author: Khalil
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB
func GormMysql() *gorm.DB {
	m := global.BB_CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", m.Username, m.Password, m.Path, m.Dbname, m.Config)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.BB_LOG.Error("Mysql 启动错误", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		return db
	}
}

//@author: Khalil
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config
func gormConfig(mod bool) *gorm.Config {
	switch global.BB_CONFIG.Mysql.LogZap {
	case "Slient":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Slient),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Error":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Error),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Warn":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Warn),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Info":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	default:
		if mod {
			return &gorm.Config{
				Logger:                                   Default.LogMode(logger.Info),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		} else {
			return &gorm.Config{
				Logger:                                   Default.LogMode(logger.Slient),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		}
	}
}
