package utils

var (
	LoginVerify          = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	IdVerify             = Rules{"ID": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ArticleVerify        = Rules{"Category": {NotEmpty()}, "Type": {Le("1")}, "Title": {NotEmpty()}, "Content": {NotEmpty()}, "AllowComment":{Le("1")}, "IsPublic": {NotEmpty()}}

)


