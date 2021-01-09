package response

type ResCategory struct {
	Id uint `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	DisplayName string `json:"display_name, omitempty"`
	SeoDesc string `json:"seo_desc, omitempty"`
}
