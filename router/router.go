/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("路由已加载")
	// 初始化
	router := gin.Default()
	// 跨域处理
	router.Use(middleware.Cros())
	// 登录验证中间件
	router.Use(middleware.CheckLogin())

	// 设置静态资源路由
	router.Static("/resource", "./public/resource")
	router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")

	// 登录注册
	login := router.Group("/")
	{
		login.GET("/captcha", controller.Login.Captcha)
		login.GET("/", controller.Login.Login)
		login.POST("/login", controller.Login.Login)
		//login.GET("/index", controller.Index.Index)
		//login.Any("/updateUserInfo", controller.Index.UpdateUserInfo)
		//login.Any("/updatePwd", controller.Index.UpdatePwd)
		//login.GET("/logout", controller.Index.Logout)
	}

	// 系统主页
	index := router.Group("index")
	{
		index.GET("/menu", controller.Index.Menu)
		index.GET("/user", controller.Index.User)
	}

	/* 职级管理 */
	level := router.Group("level")
	{
		level.GET("/list", controller.Level.List)
		level.POST("/add", controller.Level.Add)
		level.PUT("/update", controller.Level.Update)
		level.DELETE("/delete/:ids", controller.Level.Delete)
		level.PUT("/status", controller.Level.Status)
		//level.GET("/getLevelList", controller.Level.GetLevelList)
	}

	// 启动
	router.Run(":8090")
}
