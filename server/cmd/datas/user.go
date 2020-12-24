package datas

import (
	"beau-blog/global"
	"beau-blog/model"
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"os"
	"time"
)

var Users = []model.SysUser{
	{BB_MODEL: global.BB_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},  UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Nickname: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/gva_header.jpg", Email: "admin@bb.com", Phone: "001",},
	{BB_MODEL: global.BB_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "khalil", Password: "3ec063004a6f31642261936a379fde3d", Nickname: "QMPlusUser", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", Email: "admin@bb.com", Phone: "002",},
}

func InitSysUser(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1,2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("sys_users表的初始数据已存在!")
			return  nil
		}
		if err := tx.Create(&Users).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
	if err != nil {
		color.Warn.Printf("[Mysql]--> sys_users 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}