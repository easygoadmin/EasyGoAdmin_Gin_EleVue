/**
 *
 * @author 半城风雨
 * @since 2021/11/12
 * @File : OperLog
 */
package controller

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/service"
	"easygoadmin/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

var OperLog = new(operLogCtl)

type operLogCtl struct{}

func (c *operLogCtl) List(ctx *gin.Context) {
	// 参数
	var req *dto.OperLogPageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用分页查询方法
	list, count, err := service.OperLog.GetList(req)
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
