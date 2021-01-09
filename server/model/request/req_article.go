package request

type ReqArticleList struct {
	ReqPageInfo
	Tag      uint   `json:"tag, omitempty"`
	Category unit   `json:"category, omitempty"`
	OrderKey string `json:"order_key, omitempty"`
	Desc     bool   `json:"desc, omitempty"`
}

type ReqArticle struct {
	Category     int    `json:"category"`
	Type         int    `json:"type"`
	Title        string `json:"title"`
	Tags         []int  `json:"tags"`
	Content      string `json:"content"`
	AllowComment int    `json:"allow_comment"`
	IsPublic     bool   `json:"is_public"`
}
