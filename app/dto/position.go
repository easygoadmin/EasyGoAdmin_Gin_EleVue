/**
 *
 * @author 半城风雨
 * @since 2021/9/10
 * @File : position
 */
package dto

// 分页查询条件
type PositionPageReq struct {
	Name  string `form:"name"`  // 岗位名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加岗位
type PositionAddReq struct {
	Name   string `form:"name" binding:"required"`
	Status int    `form:"status" binding:"required"`
	Sort   int    `form:"sort" binding:"required"`
}

// 更新岗位
type PositionUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Status int    `form:"status" binding:"required"`
	Sort   int    `form:"sort" binding:"required"`
}

// 设置状态
type PositionStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
