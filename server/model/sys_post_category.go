package model

import "beau-blog/global"

type SysPostCategory struct {
	global.BB_MODEL
	PostId     uint          `json:"post_id" gorm:"not null unique(post_cate);"`
	CategoryId uint          `json:"category_id" gorm:"not null unique(post_cate)"`
}
