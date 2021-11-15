/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : dict
 */
package dto

// 添加字典
type DictAddReq struct {
	Name string `form:"name"  binding:"required"` // 字典名称
	Code string `form:"code"  binding:"required"` // 字典值
	Sort int    `form:"sort"  binding:"required"` // 显示顺序
	Note string `form:"note"`                     // 字典备注
}

// 修改字典
type DictUpdateReq struct {
	Id   int    `form:"id" binding:"required"`    // 主键ID
	Name string `form:"name"  binding:"required"` // 字典名称
	Code string `form:"code"  binding:"required"` // 字典值
	Sort int    `form:"sort"  binding:"required"` // 显示顺序
	Note string `form:"note"`                     // 字典备注
}
