/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : dict_data
 */
package dto

// 字典项列表查询条件
type DictDataPageReq struct {
	DictId int    `form:"dictId"` // 字典ID
	Name   string `form:"name"`   // 字典项名称
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加字典项
type DictDataAddReq struct {
	Name   string `form:"name,unique" binding:"required"` // 字典项名称
	Code   string `form:"code"        binding:"required"` // 字典项值
	DictId int    `form:"dict_id"     binding:"required"` // 字典类型ID
	Note   string `form:"note"`                           // 备注
	Sort   int    `form:"sort"        binding:"required"` // 显示顺序
}

// 更新字典项
type DictDataUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name,unique" binding:"required"` // 字典项名称
	Code   string `form:"code"        binding:"required"` // 字典项值
	DictId int    `form:"dict_id"     binding:"required"` // 字典类型ID
	Note   string `form:"note"`                           // 备注
	Sort   int    `form:"sort"        binding:"required"` // 显示顺序
}
