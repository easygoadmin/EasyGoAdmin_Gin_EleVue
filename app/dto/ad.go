// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
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
	Title       string `form:"title" binding:"required"`       // 广告标题
	AdSortId    int    `form:"ad_sort_id" binding:"required"`  // 广告位ID
	Cover       string `form:"cover"`                          // 广告图片
	Type        int    `form:"type" binding:"required"`        // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" binding:"required"` // 广告描述
	Content     string `form:"content"`                        // 广告内容
	Url         string `form:"url" binding:"required"`         // 广告链接
	Width       int    `form:"width"`                          // 广告宽度
	Height      int    `form:"height"`                         // 广告高度
	StartTime   int64  `form:"start_time" binding:"required"`  // 开始时间
	EndTime     int64  `form:"end_time" binding:"required"`    // 结束时间
	Status      int    `form:"status" binding:"required"`      // 状态：1在用 2停用
	Sort        int    `form:"sort"`                           // 排序
}

// 更新广告
type AdUpdateReq struct {
	Id          int    `form:"id" binding:"required"`
	Title       string `form:"title" binding:"required"`       // 广告标题
	AdSortId    int    `form:"ad_sort_id" binding:"required"`  // 广告位ID
	Cover       string `form:"cover"`                          // 广告图片
	Type        int    `form:"type" binding:"required"`        // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" binding:"required"` // 广告描述
	Content     string `form:"content"`                        // 广告内容
	Url         string `form:"url" binding:"required"`         // 广告链接
	Width       int    `form:"width"`                          // 广告宽度
	Height      int    `form:"height"`                         // 广告高度
	StartTime   int64  `form:"start_time" binding:"required"`  // 开始时间
	EndTime     int64  `form:"end_time" binding:"required"`    // 结束时间
	Status      int    `form:"status" binding:"required"`      // 状态：1在用 2停用
	Sort        int    `form:"sort"`                           // 排序
}

// 设置状态
type AdStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status" binding:"required"`
}
