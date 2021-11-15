/**
 *
 * @author 半城风雨
 * @since 2021/9/13
 * @File : dept
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

var Dept = new(deptCtl)

type deptCtl struct{}

func (c *deptCtl) List(ctx *gin.Context) {
	//// 参数绑定
	//var req *dto.DeptPageReq
	//if err := ctx.ShouldBind(&req); err != nil {
	//	ctx.JSON(http.StatusOK, common.JsonResult{
	//		Code: -1,
	//		Msg:  err.Error(),
	//	})
	//	return
	//}

	// 调用查询列表方法
	list, err := service.Dept.GetList(nil)
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
		Msg:  "操作成功",
		Data: list,
	})
}

func (c *deptCtl) Add(ctx *gin.Context) {
	// 参数绑定
	var req *dto.DeptAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.Dept.Add(req, utils.Uid(ctx))
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

func (c *deptCtl) Update(ctx *gin.Context) {
	// 参数绑定
	var req *dto.DeptUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.Dept.Update(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *deptCtl) Delete(ctx *gin.Context) {
	// 记录ID
	ids := ctx.Param("ids")
	// 调用删除方法
	rows, err := service.Dept.Delete(ids)
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

func (c *deptCtl) GetDeptList(ctx *gin.Context) {
	// 查询部门列表
	list := make([]model.Dept, 0)
	utils.XormDb.Where("mark=1").Asc("sort").Find(&list)
	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}
