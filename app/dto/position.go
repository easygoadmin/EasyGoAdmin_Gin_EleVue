/**
 *
 * @author 摆渡人
 * @since 2021/9/10
 * @File : position
 */
package dto

// 列表查询条件
type PositionPageReq struct {
	Name  string `form:"name"`  // 岗位名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

type PositionAddReq struct {
	Name   string `form:"name" binding:"required#岗位名称不能为空"`
	Status int    `form:sort binding:"required#岗位状态不能为空"`
	Sort   int    `form:sort binding:"required#岗位排序不能为空"`
}

type PositionUpdateReq struct {
	Id     int    `form:id binding:"required#主键ID不能为空"`
	Name   string `form:"name" binding:"required#岗位名称不能为空"`
	Status int    `form:status binding:"required#岗位状态不能为空"`
	Sort   int    `form:sort binding:"required#岗位排序不能为空"`
}

// 设置状态
type PositionStatusReq struct {
	Id     int `form:"id" binding:"required#主键ID不能为空"`
	Status int `form:"status"    binding:"required#状态不能为空"`
}
