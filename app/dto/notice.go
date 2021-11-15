/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : notice
 */
package dto

// 分页查询
type NoticePageReq struct {
	Title  string `form:"title"`  // 通知标题
	Source int    `form:"source"` // 通知来源
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加通知公告
type NoticeAddReq struct {
	Title   string `form:"title"       binding:"required"` // 通知标题
	Content string `form:"content"     binding:"required"` // 通知内容
	Source  int    `form:"source"      binding:"required"` // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"is_top"      binding:"required"` // 是否置顶：1是 2否
	Status  int    `form:"status"      binding:"required"` // 状态：1已发布 2待发布
}

// 更新通知公告
type NoticeUpdateReq struct {
	Id      int    `form:"id"		   binding:"required"`
	Title   string `form:"title"       binding:"required"` // 通知标题
	Content string `form:"content"     binding:"required"` // 通知内容
	Source  int    `form:"source"      binding:"required"` // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"is_top"      binding:"required"` // 是否置顶：1是 2否
	Status  int    `form:"status"      binding:"required"` // 状态：1已发布 2待发布
}

// 设置状态
type NoticeStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
