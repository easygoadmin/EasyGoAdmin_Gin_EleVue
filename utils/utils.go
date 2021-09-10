/**
 *
 * @author 摆渡人
 * @since 2021/8/25
 * @File : utils
 */
package utils

import (
	"easygoadmin/library/cfg"
	"easygoadmin/utils/gmd5"
	"easygoadmin/utils/gstr"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 调试模式
func AppDebug() bool {
	return false
}

// 登录用户ID
func Uid(ctx *gin.Context) int {
	// 从请求头中获取Token
	token := ctx.GetHeader("Authorization")
	// 字符串替换
	token = gstr.Replace(token, "Bearer ", "")
	claim, err := ParseToken(token)
	if err != nil {
		fmt.Println("解析token出现错误：", err)
	} else if time.Now().Unix() > claim.ExpiresAt {
		fmt.Println("时间超时")
	} else {
		//fmt.Println("username:", claim.UserId)
		//fmt.Println("username:", claim.Username)
		//fmt.Println("password:", claim.Password)
	}
	// 查询用户信息
	return claim.UserId
}

func Md5(password string) (string, error) {
	// 第一次MD5加密
	password, err := gmd5.Encrypt(password)
	if err != nil {
		return "", err
	}
	// 第二次MD5加密
	password2, err := gmd5.Encrypt(password)
	if err != nil {
		return "", err
	}
	return password2, nil
}

//获取客户端IP
func GetClientIp(ctx *gin.Context) string {
	ip := ctx.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = ctx.ClientIP()
	}
	return ip
}

func ImageUrl() string {
	// 获取配置实例
	config := cfg.Instance()
	return config.EasyGoAdmin.Image
}

// 获取文件地址
func GetImageUrl(path string) string {
	return ImageUrl() + path
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
