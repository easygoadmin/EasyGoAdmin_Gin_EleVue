/**
 *
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
)

func init() {
	fmt.Println("路由已加载")
	// 初始化
	router := gin.Default()
	// 跨域处理(要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404)
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

	/* 用户管理 */
	user := router.Group("user")
	{
		user.GET("/list", controller.User.List)
		user.GET("/detail", controller.User.Detail)
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

	// 启动
	router.Run(":8090")
}
