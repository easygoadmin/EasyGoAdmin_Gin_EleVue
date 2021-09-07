/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"easygoadmin/app/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("路由已加载")
	// 初始化
	r := gin.Default()

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
}
