/**
 *
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
