package v1

import (
	"beau-blog/global"
	"beau-blog/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var captchaStore = base64Captcha.DefaultMemStore

// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]

func Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.BB_CONFIG.Captcha.ImgHeight, global.BB_CONFIG.Captcha.ImgWidth, global.BB_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captchaStore)
	if id, b64s, err := cp.Generate(); err != nil {
		global.BB_LOG.Error("验证码获取失败", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.CaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
