/**
 *
 * @author 摆渡人
 * @since 2021/9/6
 * @File : level
 */
package dto

// 分页查询
type LevelPageReq struct {
	Name  string `form:"name"`  // 职级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加职级
type LevelAddReq struct {
	Name   string `form:"name"  binding:"required"`
	Status int    `form:"status"    binding:"required"`
	Sort   int    `form:"sort"  binding:"required"`
}

// 编辑职级
type LevelUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name"  binding:"required"`
	Status int    `form:"status"    binding:"required"`
	Sort   int    `form:"sort"  binding:"required"`
}

// 设置状态
type LevelStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
