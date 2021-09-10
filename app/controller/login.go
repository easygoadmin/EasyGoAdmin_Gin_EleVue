/**
 *
 * @author 摆渡人
 * @since 2021/9/7
 * @File : login
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/service"
	"easygoadmin/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// 用户控制器管理对象
var Login = new(loginCtl)

type loginCtl struct{}

// 系统登录
func (c *loginCtl) Login(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		var req *dto.LoginReq
		// 获取参数并验证
		if err := ctx.ShouldBind(&req); err != nil {
			// 返回错误信息
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		// 校验验证码
		verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		if !verifyRes {
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  "验证码不正确",
			})
			return
		}

		// 系统登录
		if token, err := service.Login.UserLogin(req.UserName, req.Password, ctx); err != nil {
			// 登录错误
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		} else {
			// 登录成功
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: 0,
				Msg:  "登录成功",
				Data: gin.H{
					"access_token": token,
				},
			})
		}

	}
}

// 验证码
func (c *loginCtl) Captcha(ctx *gin.Context) {
	// 验证码参数配置：字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	///create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	// 返回结果集
	ctx.JSON(http.StatusOK, common.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}
