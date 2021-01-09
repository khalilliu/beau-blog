package request

type ReqPageInfo struct {
	Page     int `json:"page" form: "page"`
	PageSize int `json:"page_size" form:"pageSize"`
}


type ReqEmpty struct {

}