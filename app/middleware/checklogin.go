// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 登录验证中间件
 * @author 半城风雨
 * @since 2021/8/20
 * @File : checkauth
 */
package middleware

import (
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gstr"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("登录验证中间件")
		// 放行设置
		urlItem := []string{"/captcha", "/login"}
		if !utils.InStringArray(ctx.Request.RequestURI, urlItem) {
			// 从请求头中获取Token
			token := ctx.GetHeader("Authorization")
			// 字符串替换
			token = gstr.Replace(token, "Bearer ", "")
			claim, err := utils.ParseToken(token)
			if err != nil {
				fmt.Println("解析token出现错误：", err)
				ctx.JSON(http.StatusOK, common.JsonResult{
					Code: 401,
					Msg:  "Token已过期",
				})
				ctx.Abort()
				return
			} else if time.Now().Unix() > claim.ExpiresAt {
				fmt.Println("时间超时")
				ctx.JSON(http.StatusOK, common.JsonResult{
					Code: 401,
					Msg:  "时间超时",
				})
				ctx.Abort()
				return
			}
		}
		// 前置中间件
		ctx.Next()
	}
}
