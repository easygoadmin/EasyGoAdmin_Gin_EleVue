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
 * 角色菜单-控制器
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
