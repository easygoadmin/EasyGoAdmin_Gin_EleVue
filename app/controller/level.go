/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : level
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model/level"
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Level = new(levelCtl)

type levelCtl struct{}

func (c *levelCtl) Index(ctx *gin.Context) {
	// 渲染模板
	response.BuildTpl(ctx, "level/index.html").WriteTpl()
}

func (c *levelCtl) List(ctx *gin.Context) {
	// 请求参数

	var req *dto.LevelPageReq
	// 请求验证
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用获取列表函数
	list, count, err := service.Level.GetList(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果集
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "操作成功",
		Count: count,
	})
}

func (c *levelCtl) Edit(ctx *gin.Context) {
	// 查询记录
	id := gconv.Int(ctx.Query("id"))
	if id > 0 {
		info := &level.Level{Id: id}
		isOk, err := info.FindOne()
		if !isOk || err != nil {
			ctx.JSON(http.StatusOK, common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 渲染模板
		response.BuildTpl(ctx, "level/edit.html").WriteTpl(gin.H{
			"info": info,
		})
	} else {
		// 渲染模板
		response.BuildTpl(ctx, "level/edit.html").WriteTpl()
	}

}

func (c *levelCtl) Add(ctx *gin.Context) {
	var req *dto.LevelAddReq
	// 请求验证
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	id, err := service.Level.Add(req, utils.Uid(ctx))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 保存成功
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}

func (c *levelCtl) Update(ctx *gin.Context) {
	// 参数验证
	var req *dto.LevelUpdateReq
	// 请求验证
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	fmt.Println(req)
	// 调用更新方法
	rows, err := service.Level.Update(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 更新成功
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *levelCtl) Delete(ctx *gin.Context) {
	var req *dto.LevelDeleteReq
	// 请求验证
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 调用删除方法
	rows, err := service.Level.Delete(req.Ids)
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 删除成功
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (c *levelCtl) Status(ctx *gin.Context) {
	var req *dto.LevelStatusReq
	// 请求验证
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	rows, err := service.Level.Status(req, utils.Uid(ctx))
	if err != nil || rows == 0 {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 设置成功
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}
