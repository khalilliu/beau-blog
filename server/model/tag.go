package model

import "beau-blog/global"

type Tag struct {
	global.BB_MODEL
	Name        string `json:"name" gorm:"not null unique; comment:标签"`
	DisplayName string `json:"display_name" gorm:"not null index VARCHAR(255); comment:标签别名"`
	SeoDesc     string `json:"seo_desc" gorm:"VARCHAR(255); comment:seo描述"`
	Num         int    `json:"num" gorm:"not null default 0; comment:文章数"`
}
