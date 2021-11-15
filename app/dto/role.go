/**
 *
 * @author 半城风雨
 * @since 2021/9/13
 * @File : role
 */
package dto

// 分页查询条件
type RolePageReq struct {
	Name  string `form:"name"`  // 角色名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加角色
type RoleAddReq struct {
	Name   string `form:"name" binding:"required"`
	Code   string `form:"code" binding:"required"`
	Status int    `form:"status" binding:"required"`
	Sort   int    `form:"sort"`
	Note   string `form:"note"`
}

// 更新角色
type RoleUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Code   string `form:"code" binding:"required"`
	Status int    `form:"status" binding:"required"`
	Sort   int    `form:"sort"`
	Note   string `form:"note"`
}

// 设置状态
type RoleStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
