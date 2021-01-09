package response

type ReqSystem struct {
	Title        string `json:"title, omitempty"`
	Keywords     string `json:"keywords, omitempty"`
	Theme        int    `json:"theme, omitempty"`
	Description  string `json:"description, omitempty"`
	RecordNumber string `json:"record_number, omitempty"`
}
