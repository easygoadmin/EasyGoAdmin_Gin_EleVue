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
 * 系统工具类
 * @author 半城风雨
 * @since 2021/8/25
 * @File : utils
 */
package utils

import (
	"easygoadmin/library/cfg"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gmd5"
	"easygoadmin/utils/gstr"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"time"
)

// 调试模式
func AppDebug() bool {
	// 获取配置实例
	config := cfg.Instance()
	return config.EasyGoAdmin.Debug
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

// 获取数据库表
func GetDatabase() (string, error) {
	config := cfg.Instance()
	if config == nil {
		fmt.Printf("参数错误")
	}

	// 获取数据库连接
	link := config.Database.Master
	if link == "" {
		return "", errors.New("数据库配置读取错误")
	}
	// 分裂字符串
	linkArr := strings.Split(link, "/")
	return strings.Split(linkArr[1], "?")[0], nil
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

// 数组反转
func Reverse(arr *[]string) {
	length := len(*arr)
	var temp string
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
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

// 判断元素是否在数组中
func InArray(value string, array []interface{}) bool {
	for _, v := range array {
		if gconv.String(v) == value {
			return true
		}
	}
	return false
}

// 附件目录
func UploadPath() string {

	// 获取配置实例
	config := cfg.Instance()
	// 附件存储路径
	upload_dir := config.EasyGoAdmin.Uploads
	if upload_dir != "" {
		return upload_dir
	} else {
		// 获取项目根目录
		curDir, _ := os.Getwd()
		return curDir + "/public/uploads"
	}
}

// 临时目录
func TempPath() string {
	return UploadPath() + "/temp"
}

// 图片存放目录
func ImagePath() string {
	return UploadPath() + "/images"
}

// 文件目录(非图片目录)
func FilePath() string {
	return UploadPath() + "/file"
}

// 创建文件夹并设置权限
func CreateDir(path string) bool {
	// 判断文件夹是否存在
	if IsExist(path) {
		return true
	}
	// 创建多层级目录
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// 判断文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		if os.IsExist(err) {
			// 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

func SaveImage(url string, dirname string) (string, error) {
	// 判断文件地址是否为空
	if gstr.Equal(url, "") {
		return "", errors.New("文件地址不能为空")
	}

	// 判断是否本站图片
	if gstr.Contains(url, ImageUrl()) {
		// 本站图片

		// 是否临时图片
		if gstr.Contains(url, "temp") {
			// 临时图片

			// 创建目录
			dirPath := ImagePath() + "/" + dirname + "/" + time.Now().Format("Ymd")
			if !CreateDir(dirPath) {
				return "", errors.New("文件目录创建失败")
			}
			// 原始图片地址
			oldPath := gstr.Replace(url, ImageUrl(), UploadPath())
			// 目标目录地址
			newPath := ImagePath() + "/" + dirname + gstr.Replace(url, ImageUrl()+"/temp", "")
			// 移动文件
			os.Rename(oldPath, newPath)
			return gstr.Replace(newPath, UploadPath(), ""), nil
		} else {
			// 非临时图片
			path := gstr.Replace(url, ImageUrl(), "")
			return path, nil
		}
	} else {
		// 远程图片
		// TODO...
	}
	return "", errors.New("保存文件异常")
}
