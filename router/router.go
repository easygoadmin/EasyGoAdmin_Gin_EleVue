/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/widget"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func init() {
	fmt.Println("路由已加载")
	// 初始化
	r := gin.Default()

	// 设置模板函数
	r.SetFuncMap(template.FuncMap{
		"widget":       widget.Widget,
		"query":        widget.Query,
		"add":          widget.Add,
		"edit":         widget.Edit,
		"delete":       widget.Delete,
		"dall":         widget.Dall,
		"expand":       widget.Expand,
		"collapse":     widget.Collapse,
		"addz":         widget.Addz,
		"switch":       widget.Switch,
		"select":       widget.Select,
		"submit":       widget.Submit,
		"icon":         widget.Icon,
		"transfer":     widget.Transfer,
		"upload_image": widget.UploadImage,
		"album":        widget.Album,
		"item":         widget.Item,
		"kindeditor":   widget.Kindeditor,
		"date":         widget.Date,
		"checkbox":     widget.Checkbox,
		"radio":        widget.Radio,
		"city":         widget.City,
	})

	// 设置静态资源路由
	r.Static("/resource", "./public/resource")
	r.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")

	// 指定模板加载目录
	r.LoadHTMLGlob("views/**/*")

	// 职级管理
	level := r.Group("/level")
	{
		level.GET("/index", controller.Level.Index)
		level.POST("/list", controller.Level.List)
		level.GET("/edit", controller.Level.Edit)
		level.POST("/add", controller.Level.Add)
		level.POST("/update", controller.Level.Update)
		level.POST("/delete", controller.Level.Delete)
		level.POST("/setStatus", controller.Level.Status)
	}

	// 启动
	r.Run()

	//// 职级管理
	//level := router.Group("/level")
	//{
	//	level.GET("/index", controller.Level.Index)
	//}

	//r := gin.New()
	//r.Use(gin.Recovery())

	//// 职级路由
	//r.Group("level", func(context *gin.Context) {
	//	r.GET("/index", controller.Level.Index)
	//})

	//r := gin.Default()
	//r.GET("/level/index", controller.Level.Index)

	//// 路由设置
	//auth := r.Group("/")
	//r.Use(middleware.CheckLogin())
	//{
	//	// 用户模块的路由接口
	//	auth.GET("/level/index", controller.Level.Index)
	//}

}
