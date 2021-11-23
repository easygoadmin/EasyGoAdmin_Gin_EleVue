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
 * 广告位Dto
 * @author 半城风雨
 * @since 2021/11/15
 * @File : ad_sort
 */
package dto

// 列表查询条件
type AdSortPageReq struct {
	Description string `form:"name"`  // 广告位描述
	Page        int    `form:"page"`  // 页码
	Limit       int    `form:"limit"` // 每页数
}

// 添加广告位
type AdSortAddReq struct {
	Description string `form:"description" binding:"required"` // 广告位描述
	ItemId      int    `form:"item_id"     binding:"required"` // 站点ID
	CateId      int    `form:"cate_id"     binding:"required"` // 栏目ID
	LocId       int    `form:"loc_id"      binding:"required"` // 广告页面位置
	Platform    int    `form:"platform"    binding:"required"` // 站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端
	Sort        int    `form:"sort"`                           // 广告位排序
}

// 更新广告位
type AdSortUpdateReq struct {
	Id          int    `form:"id" binding:"required"`
	Description string `form:"description" binding:"required"` // 广告位描述
	ItemId      int    `form:"item_id"     binding:"required"` // 站点ID
	CateId      int    `form:"cate_id"     binding:"required"` // 栏目ID
	LocId       int    `form:"loc_id"      binding:"required"` // 广告页面位置
	Platform    int    `form:"platform"    binding:"required"` // 站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端
	Sort        int    `form:"sort"`                           // 广告位排序
}
