/**
 *
 * @author 半城风雨
 * @since 2021/11/18
 * @File : upload
 */
package controller

import (
	"easygoadmin/app/service"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 控制器管理对象
var Upload = new(uploadCtl)

type uploadCtl struct{}

func (c *uploadCtl) UploadImage(ctx *gin.Context) {
	// 调用上传方法
	result, err := service.Upload.UploadImage(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 拼接图片地址
	result.FileUrl = utils.GetImageUrl(result.FileUrl)

	// 返回结果
	ctx.JSON(http.StatusOK, common.JsonResult{
		Code: 0,
		Msg:  "上传成功",
		Data: result,
	})
}
