package model

import "beau-blog/global"

type Category struct {
	global.BB_MODEL
	Name        string `json:"name" gorm:"not null unique VARCHAR(255); comment:分类名称"`
	DisplayName string `json:"display_name" gorm:"not null index unique VARCHAR(255); comment:分类别名"`
	SeoDesc     string `json:"seo_desc" gorm:"not null index VARCHAR(255); comment:seo描述"`
	ParentId    int    `json:"parent_id" gorm:"not null default 0 index INT(11); comment('父类ID')"`
}
