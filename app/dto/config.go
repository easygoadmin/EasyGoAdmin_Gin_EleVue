/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : config
 */
package dto

// 添加配置
type ConfigAddReq struct {
	Name string `form:"name"  binding:"required"` // 配置名称
	Sort int    `form:"sort"  binding:"required"` // 显示顺序
}

// 修改配置
type ConfigUpdateReq struct {
	Id   int    `form:"id" binding:"required"`    // 主键ID
	Name string `form:"name"  binding:"required"` // 配置名称
	Sort int    `form:"sort"  binding:"required"` // 显示顺序
}
