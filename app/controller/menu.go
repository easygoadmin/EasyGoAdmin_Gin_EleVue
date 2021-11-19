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
 * 菜单管理-控制器
 * @author 半城风雨
 * @since 2021/9/13
 * @File : menu
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Menu = new(menuCtl)

type menuCtl struct{}

func (c *menuCtl) List(ctx *gin.Context) {
	//// 参数绑定
	//var req *dto.MenuQueryReq
	//if err := ctx.Bind(&req); err != nil {
	//	ctx.JSON(http.StatusOK, common.JsonResult{
	//		Code: -1,
	//		Msg:  err.Error(),
	//	})
	//	return
	//}

	// 调用查询列表方法
	list, err := service.Menu.GetList(nil)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}

func (c *menuCtl) Detail(ctx *gin.Context) {
	// 记录ID
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}
	info := model.Menu{Id: gconv.Int(id)}
	has, err := info.Get()
	if err != nil || !has {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 菜单信息
	menu := dto.MenuInfoVo{}
	menu.Menu = info
	// 获取权限节点
	if info.Type == 0 {
		// 获取角色菜单权限列表
		var menuList []model.Menu
		utils.XormDb.Where("parent_id=?", info.Id).Where("type=1").Where("mark=1").Find(&menuList)
		checkedList := make([]int, 0)
		for _, v := range menuList {
			checkedList = append(checkedList, v.Sort)
		}
		if len(checkedList) > 0 {
			menu.CheckedList = checkedList
		} else {
			menu.CheckedList = make([]int, 0)
		}
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: menu,
	})
}

func (c *menuCtl) Add(ctx *gin.Context) {
	// 参数绑定
	var req *dto.MenuAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.Menu.Add(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *menuCtl) Update(ctx *gin.Context) {
	// 参数绑定
	var req *dto.MenuUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.Menu.Update(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *menuCtl) Delete(ctx *gin.Context) {
	// 记录ID
	ids := ctx.Param("ids")
	// 调用删除方法
	rows, err := service.Menu.Delete(ids)
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}
