package model

import (
	"beau-blog/global"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.BB_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username  string    `json:"username" gorm:"comment:用户名"`
	Password  string    `json:"password" gorm:"comment:用户登录密码"`
	Nickname  string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg string    `json:"header_img" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Email     string    `json:"email" gorm:"default:'';comment:邮箱"`
	Phone     string    `json:"phone" gorm:"default:'';comment:手机号"`
}
