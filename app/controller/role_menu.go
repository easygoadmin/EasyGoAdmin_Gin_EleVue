/**
 *
 * @author 半城风雨
 * @since 2021/9/14
 * @File : role_menu
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/service"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

var RoleMenu = new(roleMenuCtl)

type roleMenuCtl struct{}

func (c *roleMenuCtl) Index(ctx *gin.Context) {
	// 角色ID
	roleId := ctx.Param("roleId")
	if roleId == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "角色ID不能为空",
		})
		return
	}

	// 获取角色菜单权限列表
	list, err := service.RoleMenu.GetRoleMenuList(gconv.Int(roleId))
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
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *roleMenuCtl) Save(ctx *gin.Context) {
	var req *dto.RoleMenuSaveReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用保存方法
	err := service.RoleMenu.Save(req)
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
		Msg:  "保存成功",
	})
}
