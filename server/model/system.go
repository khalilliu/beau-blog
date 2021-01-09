package model

import "beau-blog/global"

type System struct {
	global.BB_MODEL
	Theme        int    `json:"theme" gorm:"not null TINYINT(4) default 0;comment:主题"`
	Title        string `json:"title" gorm:"not null ;comment:网站title"`
	Keywords     string `json:"keywords" gorm:"not null VARCHAR(255); comment:网站关键字"`
	Description  string `json:"description" gorm:"not null; comment:网站描述"`
	RecordNumber string `json:"record_number" gorm:"not null; comment:网站备案"`
}
