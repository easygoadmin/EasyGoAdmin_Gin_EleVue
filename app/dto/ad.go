// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 广告Dto
 * @author 半城风雨
 * @since 2021/11/15
 * @File : ad
 */
package dto

// 列表查询
type AdPageReq struct {
	Title string `form:"title"` // 广告标题
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加广告
type AdAddReq struct {
	Title       string `form:"title"       binding:"required"` // 广告标题
	AdSortId    int    `form:"ad_sort_id"  binding:"required"` // 广告位ID
	Cover       string `form:"cover"`                          // 广告图片
	Type        int    `form:"type"        binding:"required"` // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" binding:"required"` // 广告描述
	Content     string `form:"content"`                        // 广告内容
	Url         string `form:"url"         binding:"required"` // 广告链接
	Width       int    `form:"width"`                          // 广告宽度
	Height      int    `form:"height"`                         // 广告高度
	StartTime   int64  `form:"start_time"  binding:"required"` // 开始时间
	EndTime     int64  `form:"end_time"    binding:"required"` // 结束时间
	Status      int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort        int    `form:"sort"`                           // 排序
}

// 更新广告
type AdUpdateReq struct {
	Id          int    `form:"id" binding:"required"`
	Title       string `form:"title"       binding:"required"` // 广告标题
	AdSortId    int    `form:"ad_sort_id"  binding:"required"` // 广告位ID
	Cover       string `form:"cover"`                          // 广告图片
	Type        int    `form:"type"        binding:"required"` // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" binding:"required"` // 广告描述
	Content     string `form:"content"`                        // 广告内容
	Url         string `form:"url"         binding:"required"` // 广告链接
	Width       int    `form:"width"`                          // 广告宽度
	Height      int    `form:"height"`                         // 广告高度
	StartTime   int64  `form:"start_time"  binding:"required"` // 开始时间
	EndTime     int64  `form:"end_time"    binding:"required"` // 结束时间
	Status      int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort        int    `form:"sort"`                           // 排序
}

// 设置状态
type AdStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
