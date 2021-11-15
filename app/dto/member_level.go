/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : member_level
 */
package dto

// 查询会员等级
type MemberLevelPageReq struct {
	Name  string `form:"name"`  // 等级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加会员等级
type MemberLevelAddReq struct {
	Name string `form:"name"        v:"name"` // 级别名称
	Sort int    `form:"sort"        v:"sort"` // 排序号
}

// 更新会员等级
type MemberLevelUpdateReq struct {
	Id   int    `form:"id" binding:"required"`
	Name string `form:"name"        binding:"required"` // 级别名称
	Sort int    `form:"sort"        binding:"required"` // 排序号
}
