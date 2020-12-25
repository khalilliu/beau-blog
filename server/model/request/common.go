package request

type PageInfo struct {
	Page     int `json:"page" form: "page"`
	PageSize int `json:"page_size" form:"pageSize"`
}


type Empty struct {

}