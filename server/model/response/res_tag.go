package response

type ResTag struct {
	Id          uint   `json:"id" json:"id, omitempty"`
	Name        string `json:"name, omitempty"`
	DisplayName string `json:"display_name, omitempty"`
	SeoDesc     string `json:"seo_desc, omitempty"`
	Num         int    `json:"num, omitempty"`
}
