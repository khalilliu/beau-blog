package model

import "beau-blog/global"

type SysPostTag struct {
	global.BB_MODEL
	PostId int      `json:"post_id" gorm:"not null unique(post_tag);"`
	TagId  int      `json:"tag_id" gorm:"not null unique(post_tag)"`
	Post   *SysPost `json:"post" gorm:"-"`
	Tag    *SysTag  `json:"tag" gorm:"-"`
}
