/**
 *
 * @author 半城风雨
 * @since 2021/11/18
 * @File : checkauth
 */
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("鉴权中间件")

		// TODO...

		// 前置中间件
		context.Next()
	}
}
