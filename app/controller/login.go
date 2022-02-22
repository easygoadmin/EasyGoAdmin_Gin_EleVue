// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 登录-控制器
 * @author 半城风雨
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
