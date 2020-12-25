package model

import "beau-blog/global"

type Post struct {
	global.BB_MODEL
	CategoryId      uint      `json:"category_id" gorm:"comment:文章类别id"`
	Category        *Category `json:"category,omitempty" gorm:"-;comment:文章类别"`
	UserId          uint      `json:"user_id"`
	Type            int       `json:"type" gorm:"not null default 0; comment:0 文章, 1 页面"`
	Status          int       `json:"status" gorm:"not null default 0; comment:0 为草稿，1 为待审核，2 为已拒绝，3 为已经发布"`
	Title           string    `json:"title" gorm:"not null VARCHAR(255); comment:文章标题"`
	Tags            string    `json:"tags" gorm:"type:text;comment:标签"`
	Abstract        string    `json:"abstract" gorm:"type:text;comment:简介"`
	MarkdownContent string    `json:"markdown_content" gorm:"not null LONGTEXT;comment:文章markdown内容"`
	Content         string    `json:"content" gorm:"not null LONGTEXT;comment:文章内容"`
	AllowComment    int       `json:"allow_comment" gorm:"not null default 1; comment: 0 不可评论 1 可评论"`
	IsPublic        bool      `json:"is_public" gorm:"not null default 1; comment: 0 不公开 1 公开"`
	CommentNum      int       `json:"comment_num" gorm:"not null default 0; comment:文章评论数"`
	ViewNum         int       `json:"view_num" gorm:"not null default 0;comment:文章查看人次"`
	Options         string    `json:"options" gorm:"not null; comment:一些结构"`
}
