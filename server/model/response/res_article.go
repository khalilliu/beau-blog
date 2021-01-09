package response

import (
	"beau-blog/model"
	"time"
)

type ResArticle struct {
	Id        uint      `json:"id, omitempty"`
	UserId    uint      `json:"user_id, omitempty"`
	Title     string    `json:"title, omitempty"`
	Abstract  string    `json:"abstract, omitempty"`
	Content   string    `json:"content, omitempty"`
	DeletedAt time.Time `json:"deleted_at, omitempty"`
	CreatedAt time.Time `json:"created_at, omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ReqArticleDetail struct {
	Article  ResArticle   `json:"post, omitempty"`
	Tags     []ResTag       `json:"tags, omitempty"`
	Category ResCategory    `json:"category, omitempty"`
	Author   ResUser        `json:"author, omitempty"`
	LastPost *model.Article `json:"last_post, omitempty"`
	NextPost *model.Article `json:"next_post, omitempty"`
}
