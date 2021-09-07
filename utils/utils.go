/**
 *
 * @author 摆渡人
 * @since 2021/8/25
 * @File : utils
 */
package utils

import "github.com/gin-gonic/gin"

// 登录用户ID
func Uid(ctx *gin.Context) int {
	return 1
}

// 调试模式
func AppDebug() bool {
	return false
}
