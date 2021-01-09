package request

type ReqRegister struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img"`
}

type ReqLogin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

type ReqChangePassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

