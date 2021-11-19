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
 * 代码生成器-控制器
 * @author 半城风雨
 * @since 2021/11/15
 * @File : generate
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

// 控制器管理对象
var Generate = new(generateCtl)

type generateCtl struct{}

func (c *generateCtl) List(ctx *gin.Context) {
	// 参数验证
	var req *dto.GeneratePageReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用查询列表方法
	list, err := service.Generate.GetList(req)
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
		Count: gconv.Int64(len(list)),
	})
}

func (c *generateCtl) Generate(ctx *gin.Context) {
	// 调用生成方法
	err := service.Generate.Generate(ctx)
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
		Msg:  "模块生成成功",
	})
}
