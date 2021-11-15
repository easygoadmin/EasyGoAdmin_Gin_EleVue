/**
 *
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
