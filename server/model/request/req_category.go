package request

type ReqCategory struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	SeoDesc     string `json:"seo_desc"`
	ParentId    int    `json:"parent_id"`
}