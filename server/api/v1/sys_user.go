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


