/**
 *
 * @author 半城风雨
 * @since 2021/9/7
 * @File : index
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户API管理对象
var Index = new(indexCtl)

type indexCtl struct{}

// 获取系统菜单
func (c *indexCtl) Menu(ctx *gin.Context) {
	// 获取菜单列表
	menuList := service.Menu.GetPermissionList(utils.Uid(ctx))
	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: menuList,
	})
}

// 获取用户信息
func (c *indexCtl) User(ctx *gin.Context) {
	// 获取用户信息
	user := model.User{}
	has, err := utils.XormDb.Id(utils.Uid(ctx)).Get(&user)
	if err != nil && !has {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 头像
	user.Avatar = utils.GetImageUrl(user.Avatar)
	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: user,
	})
}

// 个人中心
func (c *indexCtl) UpdateUserInfo(ctx *gin.Context) {
	// 参数验证
	var req *dto.UserInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 更新信息
	_, err := service.User.UpdateUserInfo(req, utils.Uid(ctx))
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
		Msg:  "更新成功",
	})
}

// 更新密码
func (c *indexCtl) UpdatePwd(ctx *gin.Context) {
	// 参数验证
	var req *dto.UpdatePwd
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新密码方法
	rows, err := service.User.UpdatePwd(req, utils.Uid(ctx))
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
		Msg:  "更新密码成功",
	})
}

// 注销系统
func (c *indexCtl) Logout(ctx *gin.Context) {
	// 注销系统并返回提示
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "注销成功",
	})
}
