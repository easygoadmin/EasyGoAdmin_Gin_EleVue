/**
 *
 * @author 半城风雨
 * @since 2021/9/7
 * @File : index
 */
package controller

import (
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
