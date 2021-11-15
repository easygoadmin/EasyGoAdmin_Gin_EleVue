/**
 *
 * @author 半城风雨
 * @since 2021/11/11
 * @File : user
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

var User = new(userCtrl)

type userCtrl struct{}

func (c *userCtrl) List(ctx *gin.Context) {
	// 参数
	var req *dto.UserPageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用获取列表方法
	list, count, err := service.User.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "查询成功",
		Count: count,
	})
}

func (c *userCtrl) Detail(ctx *gin.Context) {
	// 记录ID
	id := ctx.Param("id")
	if id != "" {
		// 编辑
		info := &model.User{Id: gconv.Int(id)}
		has, err := info.Get()
		if err != nil || !has {
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		var userInfo = vo.UserInfoVo{}
		userInfo.User = *info
		// 头像
		userInfo.Avatar = utils.GetImageUrl(info.Avatar)
		// 角色列表
		roleList := service.UserRole.GetUserRoleList(info.Id)
		if len(roleList) > 0 {
			userInfo.RoleList = roleList
		} else {
			userInfo.RoleList = make([]model.Role, 0)
		}
		// 省市区
		cityList := make([]string, 0)
		// 省份编号
		cityList = append(cityList, info.ProvinceCode)
		// 城市编号
		cityList = append(cityList, info.CityCode)
		// 县区编号
		cityList = append(cityList, info.DistrictCode)
		userInfo.City = cityList

		// 返回结果
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: 0,
			Msg:  "查询成功",
			Data: userInfo,
		})
	}
}

func (c *userCtrl) Add(ctx *gin.Context) {
	// 参数
	var req *dto.UserAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.User.Add(req, utils.Uid(ctx))
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

func (c *userCtrl) Update(ctx *gin.Context) {
	// 参数
	var req *dto.UserUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.User.Update(req, utils.Uid(ctx))
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

func (c *userCtrl) Delete(ctx *gin.Context) {
	// 参数
	ids := ctx.Param("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}

	// 调用删除方法
	rows, err := service.User.Delete(ids)
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

func (c *userCtrl) Status(ctx *gin.Context) {
	// 参数
	var req *dto.UserStatusReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用设置方法
	rows, err := service.User.Status(req, utils.Uid(ctx))
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
		Msg:  "设置成功",
	})
}

func (c *userCtrl) ResetPwd(ctx *gin.Context) {
	// 参数验证
	var req *dto.UserResetPwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用重置密码方法
	rows, err := service.User.ResetPwd(req.Id, utils.Uid(ctx))
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
		Msg:  "重置密码成功",
	})
}

func (c *userCtrl) CheckUser(ctx *gin.Context) {
	// 参数验证
	var req *dto.CheckUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用检查用户方法
	user, err := service.User.CheckUser(req)
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
		Data: user,
	})
}
