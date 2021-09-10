/**
 *
 * @author 摆渡人
 * @since 2021/9/7
 * @File : login
 */
package service

import (
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"time"
)

var x *xorm.Engine

// 中间件管理服务
var Login = new(loginService)

type loginService struct{}

// 系统登录
func (s *loginService) UserLogin(username, password string, ctx *gin.Context) (string, error) {
	// 查询用户
	var user model.User
	isOk, err := utils.XormDb.Where("username=? and mark=1", username).Get(&user)
	if err != nil && isOk {
		return "", errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.Md5(password + user.Username)
	if user.Password != pwd {
		return "", errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		return "", errors.New("您的账号已被禁用,请联系管理员")
	}

	// 更新登录时间、登录IP
	utils.XormDb.Id(user.Id).Update(&model.User{LoginTime: time.Now(), LoginIp: utils.GetClientIp(ctx), UpdateTime: time.Now()})

	// 生成Token
	token, _ := utils.GenerateToken(user.Id, user.Username, user.Password)
	fmt.Println("生成的token:", token)

	// 返回token
	return token, nil
}
