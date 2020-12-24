package datas

import (
	"beau-blog/model"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

func InitMysqlData(db *gorm.DB) {
	InitSysUser(db)
	InitSysCagegory(db)
	InitSysTag(db)
	InitSysPost(db)
}

func InitMysqlTables(db *gorm.DB) {
	var err error

	err = db.Migrate{
		model.SysUser{},
		model.SysCategory{},
		model.SysPostCategory{},
		model.SysPost{},
		model.SysPostTag{},
		model.SysTag{},
		model.SysUser{},
	}

	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(1)
	}
	color.Info.Println("[Mysql]-->初始化数据表成功")
}

