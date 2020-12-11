package request

type Register struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

type ChangePassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

