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
 * 广告管理-控制器
 * @author 半城风雨
 * @since 2021/11/15
 * @File : ad
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Ad = new(adCtl)

type adCtl struct{}

func (c *adCtl) List(ctx *gin.Context) {
	// 参数
	var req *dto.AdPageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用分页查询方法
	list, count, err := service.Ad.GetList(req)
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
		Msg:   "查询成功",
		Data:  list,
		Count: count,
	})
}

func (c *adCtl) Add(ctx *gin.Context) {
	// 参数
	var req *dto.AdAddReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用添加方法
	rows, err := service.Ad.Add(req, utils.Uid(ctx))
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

func (c *adCtl) Update(ctx *gin.Context) {
	// 参数
	var req *dto.AdUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用更新方法
	rows, err := service.Ad.Update(req, utils.Uid(ctx))
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

func (c *adCtl) Delete(ctx *gin.Context) {
	// 记录ID
	ids := ctx.Param("ids")
	if ids == "" {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
		return
	}

	// 调用删除方法
	rows, err := service.Ad.Delete(ids)
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

func (c *adCtl) Status(ctx *gin.Context) {
	// 参数绑定
	var req *dto.AdStatusReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用设置状态方法
	rows, err := service.Ad.Status(req, utils.Uid(ctx))
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
