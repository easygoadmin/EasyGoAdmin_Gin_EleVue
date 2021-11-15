/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item
 */
package dto

// 分页查询条件
type ItemPageReq struct {
	Name  string `form:"name"`  // 站点名称
	Type  int    `form:"type"`  // 站点类型
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加站点
type ItemAddReq struct {
	Name   string `form:"name"        binding:"required"` // 站点名称
	Type   int    `form:"type"        binding:"required"` // 站点类型:1普通站点 2其他
	Url    string `form:"url"         binding:"required"` // 站点地址
	Image  string `form:"image"`                          // 站点图片
	Status int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Note   string `form:"note"`                           // 站点备注
	Sort   int    `form:"sort"        binding:"required"` // 显示顺序
}

// 更新站点
type ItemUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name"        binding:"required"` // 站点名称
	Type   int    `form:"type"        binding:"required"` // 站点类型:1普通站点 2其他
	Url    string `form:"url"         binding:"required"` // 站点地址
	Image  string `form:"image"`                          // 站点图片
	Status int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Note   string `form:"note"`                           // 站点备注
	Sort   int    `form:"sort"        binding:"required"` // 显示顺序
}

// 设置状态
type ItemStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
