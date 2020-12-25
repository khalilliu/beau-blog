package model

import "beau-blog/global"

type Category struct {
	global.BB_MODEL
	Name      string `json:"name" gorm:"not null unique VARCHAR(255); comment:分类名称"`
	HeaderImg string `json:"header_img" gorm:"not null;comment:分类头图'"`
	Pid       int    `json:"pid" gorm:"not null default 0; comment:排序"`
	Intro     string `json:"intro"`
}
