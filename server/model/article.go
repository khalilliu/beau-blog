package model

import "beau-blog/global"

type Article struct {
	global.BB_MODEL
	UserId          uint      `json:"user_id" gorm:"comment:作者"`
	Type            int       `json:"type" gorm:"not null default 0; comment:0 文章, 1 页面"`
	Title           string    `json:"title" gorm:"not null VARCHAR(255); comment:文章标题"`
	Abstract        string    `json:"abstract" gorm:"type:text;comment:摘要"`
	Content         string    `json:"content" gorm:"not null LONGTEXT;comment:文章内容"`
	AllowComment    int       `json:"allow_comment" gorm:"not null default 1; comment: 0 不可评论 1 可评论"`
	IsPublic        bool      `json:"is_public" gorm:"not null default 1; comment: 0 不公开 1 公开"`
	CommentNum      int       `json:"comment_num" gorm:"not null default 0; comment:文章评论数"`
	ViewNum         int       `json:"view_num" gorm:"not null default 0;comment:文章查看人次"`
}
