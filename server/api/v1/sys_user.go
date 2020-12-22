package v1

impor (
	"github.com/gin-gonic/gin"

)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c);
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		User := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(User); err != nil {
			global.BB_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("登陆失败! 用户名不存在或者密码错误", c);
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c);
	}
}

// 登录成功后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.BB_CONFIG.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		BufferTime:  60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
		ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
		Issuer:    "bbPlus",                       // 签名的发行者
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.BB_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.BB_CONFIG.System.UseMultipoing {
		response.OkWithDetailed(response.LoginResponse{
			User: user,
			Token: token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	} 

}


// @Tags SysUser
// @Summary 用户登录
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {

}

// @Tags SysUser
// @Summary 用户修改密码
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/register [post]



// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/setUserInfo [put]
func Register(c *gin.Context) {

}


