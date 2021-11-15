/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : ad
 */
package dto

import "time"

// 列表查询
type AdPageReq struct {
	Title string `form:"title"` // 广告标题
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加广告
type AdAddReq struct {
	Title       string    `form:"title"       binding:"required"` // 广告标题
	AdSortId    int       `form:"ad_sort_id"  binding:"required"` // 广告位ID
	Cover       string    `form:"cover"`                          // 广告图片
	Type        int       `form:"type"        binding:"required"` // 广告格式：1图片 2文字 3视频 4推荐
	Description string    `form:"description" binding:"required"` // 广告描述
	Content     string    `form:"content"`                        // 广告内容
	Url         string    `form:"url"         binding:"required"` // 广告链接
	Width       int       `form:"width"`                          // 广告宽度
	Height      int       `form:"height"`                         // 广告高度
	StartTime   time.Time `form:"start_time"  binding:"required"` // 开始时间
	EndTime     time.Time `form:"end_time"    binding:"required"` // 结束时间
	Status      int       `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort        int       `form:"sort"        binding:"required"` // 排序
}

// 更新广告
type AdUpdateReq struct {
	Id          int       `form:"id" binding:"required"`
	Title       string    `form:"title"       binding:"required"` // 广告标题
	AdSortId    int       `form:"ad_sort_id"  binding:"required"` // 广告位ID
	Cover       string    `form:"cover"`                          // 广告图片
	Type        int       `form:"type"        binding:"required"` // 广告格式：1图片 2文字 3视频 4推荐
	Description string    `form:"description" binding:"required"` // 广告描述
	Content     string    `form:"content"`                        // 广告内容
	Url         string    `form:"url"         binding:"required"` // 广告链接
	Width       int       `form:"width"`                          // 广告宽度
	Height      int       `form:"height"`                         // 广告高度
	StartTime   time.Time `form:"start_time"  binding:"required"` // 开始时间
	EndTime     time.Time `form:"end_time"    binding:"required"` // 结束时间
	Status      int       `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort        int       `form:"sort"        binding:"required"` // 排序
}

// 设置状态
type AdStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
