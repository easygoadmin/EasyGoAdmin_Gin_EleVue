// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2021 EasyGoAdmin深圳研发中心
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
 * 系统路由
 * @author 半城风雨
 * @since 2021/8/20
 * @File : router
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

// 日期格式转换
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func init() {
	fmt.Println("路由已加载")
	// 初始化
	router := gin.Default()
	//// 自定义标识符，因为默认{{}}这种标识在前端框架中也有使用，会产生冲突
	//// 注意：改完标识符之后别忘了把模板里原先的标识符一起改掉
	//router.Delims("{[", "]}")

	// 自定义模板方法
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	//// 指定模板加载目录
	//router.LoadHTMLGlob("views/**")
	// 跨域处理(要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404)
	router.Use(middleware.Cros())
	// 登录验证中间件
	router.Use(middleware.CheckLogin())
	// 鉴权拦截器中间件
	router.Use(middleware.CheckAuth())

	//// 验证登录后放行路由请求
	//router.Use(middleware.CheckLogin())
	//{
	//	// 写路由
	//}

	// 设置静态资源路由
	router.Static("/resource", "./public/resource")
	router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")

	/* 文件上传 */
	upload := router.Group("upload")
	{
		// 上传图片
		upload.POST("/uploadImage", controller.Upload.UploadImage)
	}

	// 登录注册
	login := router.Group("/")
	{
		login.GET("/captcha", controller.Login.Captcha)
		login.GET("/", controller.Login.Login)
		login.POST("/login", controller.Login.Login)
		login.Any("/updateUserInfo", controller.Index.UpdateUserInfo)
		login.Any("/updatePwd", controller.Index.UpdatePwd)
		login.GET("/logout", controller.Index.Logout)
	}

	// 系统主页
	index := router.Group("index")
	{
		index.GET("/menu", controller.Index.Menu)
		index.GET("/user", controller.Index.User)
	}

	/* 用户管理 */
	user := router.Group("user")
	{
		user.GET("/list", controller.User.List)
		user.GET("/detail/:id", controller.User.Detail)
		user.POST("/add", controller.User.Add)
		user.PUT("/update", controller.User.Update)
		user.DELETE("/delete/:ids", controller.User.Delete)
		user.PUT("/status", controller.User.Status)
		user.PUT("/resetPwd", controller.User.ResetPwd)
		user.GET("/checkUser", controller.User.CheckUser)
	}

	/* 职级管理 */
	level := router.Group("level")
	{
		level.GET("/list", controller.Level.List)
		level.POST("/add", controller.Level.Add)
		level.PUT("/update", controller.Level.Update)
		level.DELETE("/delete/:ids", controller.Level.Delete)
		level.PUT("/status", controller.Level.Status)
		level.GET("/getLevelList", controller.Level.GetLevelList)
	}

	/* 岗位管理 */
	position := router.Group("position")
	{
		position.GET("/list", controller.Position.List)
		position.POST("/add", controller.Position.Add)
		position.PUT("/update", controller.Position.Update)
		position.DELETE("/delete/:ids", controller.Position.Delete)
		position.PUT("/status", controller.Position.Status)
		position.GET("/getPositionList", controller.Position.GetPositionList)
	}

	/* 部门管理 */
	dept := router.Group("dept")
	{
		dept.GET("/list", controller.Dept.List)
		dept.POST("/add", controller.Dept.Add)
		dept.PUT("/update", controller.Dept.Update)
		dept.DELETE("/delete/:ids", controller.Dept.Delete)
		dept.GET("/getDeptList", controller.Dept.GetDeptList)
	}

	/* 菜单管理 */
	menu := router.Group("menu")
	{
		menu.GET("/list", controller.Menu.List)
		menu.GET("/detail", controller.Menu.Detail)
		menu.POST("/add", controller.Menu.Add)
		menu.PUT("/update", controller.Menu.Update)
		menu.DELETE("/delete/:ids", controller.Menu.Delete)
	}

	/* 角色路由 */
	role := router.Group("role")
	{
		role.GET("/list", controller.Role.List)
		role.POST("/add", controller.Role.Add)
		role.PUT("/update", controller.Role.Update)
		role.DELETE("/delete/:ids", controller.Role.Delete)
		role.PUT("/status", controller.Role.Status)
		role.GET("/getRoleList", controller.Role.GetRoleList)
	}

	/* 角色菜单权限 */
	roleMenu := router.Group("rolemenu")
	{
		roleMenu.GET("/index/:roleId", controller.RoleMenu.Index)
		roleMenu.POST("/save", controller.RoleMenu.Save)
	}

	/* 登录日志 */
	loginLog := router.Group("loginlog")
	{
		loginLog.GET("/list", controller.LoginLog.List)
		loginLog.DELETE("/delete/:ids", controller.LoginLog.Delete)
	}

	/* 操作日志 */
	operLog := router.Group("operlog")
	{
		operLog.GET("/list", controller.OperLog.List)
	}

	/* 字典管理 */
	dict := router.Group("dict")
	{
		dict.GET("/list", controller.Dict.List)
		dict.POST("/add", controller.Dict.Add)
		dict.PUT("/update", controller.Dict.Update)
		dict.DELETE("/delete/:ids", controller.Dict.Delete)
	}

	/* 字典项管理 */
	dictdata := router.Group("dictdata")
	{
		dictdata.GET("/list", controller.DictData.List)
		dictdata.POST("/add", controller.DictData.Add)
		dictdata.PUT("/update", controller.DictData.Update)
		dictdata.DELETE("/delete/:ids", controller.DictData.Delete)
	}

	/* 配置管理 */
	config := router.Group("config")
	{
		config.GET("/list", controller.Config.List)
		config.POST("/add", controller.Config.Add)
		config.PUT("/update", controller.Config.Update)
		config.DELETE("/delete/:ids", controller.Config.Delete)
	}

	/* 配置项管理 */
	configdata := router.Group("configdata")
	{
		configdata.GET("/list", controller.ConfigData.List)
		configdata.POST("/add", controller.ConfigData.Add)
		configdata.PUT("/update", controller.ConfigData.Update)
		configdata.DELETE("/delete/:ids", controller.ConfigData.Delete)
		configdata.PUT("/status", controller.ConfigData.Status)
	}

	/* 通知管理 */
	notice := router.Group("notice")
	{
		notice.GET("/list", controller.Notice.List)
		notice.POST("/add", controller.Notice.Add)
		notice.PUT("/update", controller.Notice.Update)
		notice.DELETE("/delete/:ids", controller.Notice.Delete)
		notice.PUT("/status", controller.Notice.Status)
	}

	/* 城市管理 */
	city := router.Group("city")
	{
		city.GET("/list", controller.City.List)
		city.POST("/add", controller.City.Add)
		city.PUT("/update", controller.City.Update)
		city.DELETE("/delete/:ids", controller.City.Delete)
		city.POST("/getChilds", controller.City.GetChilds)
	}

	/* 友链管理 */
	link := router.Group("link")
	{
		link.GET("/list", controller.Link.List)
		link.POST("/add", controller.Link.Add)
		link.PUT("/update", controller.Link.Update)
		link.DELETE("/delete/:ids", controller.Link.Delete)
		link.PUT("/status", controller.Link.Status)
	}

	/* 站点管理 */
	item := router.Group("item")
	{
		item.GET("/list", controller.Item.List)
		item.POST("/add", controller.Item.Add)
		item.PUT("/update", controller.Item.Update)
		item.DELETE("/delete/:ids", controller.Item.Delete)
		item.PUT("/status", controller.Item.Status)
		item.GET("/getItemList", controller.Item.GetItemList)
	}

	/* 栏目管理 */
	itemcate := router.Group("itemcate")
	{
		itemcate.GET("/list", controller.ItemCate.List)
		itemcate.POST("/add", controller.ItemCate.Add)
		itemcate.PUT("/update", controller.ItemCate.Update)
		itemcate.DELETE("/delete/:ids", controller.ItemCate.Delete)
		itemcate.GET("/getCateList", controller.ItemCate.GetCateList)
	}

	/* 广告位管理 */
	adsort := router.Group("adsort")
	{
		adsort.GET("/list", controller.AdSort.List)
		adsort.POST("/add", controller.AdSort.Add)
		adsort.PUT("/update", controller.AdSort.Update)
		adsort.DELETE("/delete/:ids", controller.AdSort.Delete)
		adsort.GET("/getAdSortList", controller.AdSort.GetAdSortList)
	}

	/* 广告管理 */
	ad := router.Group("ad")
	{
		ad.GET("/list", controller.Ad.List)
		ad.POST("/add", controller.Ad.Add)
		ad.PUT("/update", controller.Ad.Update)
		ad.DELETE("/delete/:ids", controller.Ad.Delete)
		ad.PUT("/status", controller.Ad.Status)
	}

	/* 会员等级 */
	memberlevel := router.Group("memberlevel")
	{
		memberlevel.GET("/list", controller.MemberLevel.List)
		memberlevel.POST("/add", controller.MemberLevel.Add)
		memberlevel.PUT("/update", controller.MemberLevel.Update)
		memberlevel.DELETE("/delete/:ids", controller.MemberLevel.Delete)
		memberlevel.GET("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	}

	/* 会员管理 */
	member := router.Group("member")
	{
		member.GET("/list", controller.Member.List)
		member.POST("/add", controller.Member.Add)
		member.PUT("/update", controller.Member.Update)
		member.DELETE("/delete/:ids", controller.Member.Delete)
		member.PUT("/status", controller.Member.Status)
	}

	/* 网站设置 */
	configweb := router.Group("configweb")
	{
		configweb.GET("/index", controller.ConfigWeb.Index)
		configweb.PUT("/save", controller.ConfigWeb.Save)
	}

	/* 代码生成器 */
	generate := router.Group("generate")
	{
		generate.GET("/list", controller.Generate.List)
		generate.POST("/generate", controller.Generate.Generate)
	}

	/* 演示一 */
	example := router.Group("example")
	{
		example.GET("/list", controller.Example.List)
		example.POST("/add", controller.Example.Add)
		example.PUT("/update", controller.Example.Update)
		example.DELETE("/delete/:ids", controller.Example.Delete)

		example.PUT("/status", controller.Example.Status)
		example.PUT("/isVip", controller.Example.IsVip)
	}

	// 启动
	router.Run(":9090")
}
