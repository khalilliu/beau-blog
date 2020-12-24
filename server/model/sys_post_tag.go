package model

import "beau-blog/global"

type SysPostTag struct {
	global.BB_MODEL
	PostId uint      `json:"post_id" gorm:"not null unique(post_tag);"`
	TagId  uint      `json:"tag_id" gorm:"not null unique(post_tag)"`
}
