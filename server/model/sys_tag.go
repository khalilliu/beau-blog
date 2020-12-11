package model

import "beau-blog/global"

type SysTag struct {
	global.BB_MODEL
	Name  string `json:"name" gorm:"not null unique; comment:标签"`
	Intro string `json:"intro" gorm:"not null;comment:标签介绍"`
}
