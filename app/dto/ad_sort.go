/**
 *
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
	Sort        int    `form:"sort"        binding:"required"` // 广告位排序
}

// 更新广告位
type AdSortUpdateReq struct {
	Id          int    `form:"id" binding:"required"`
	Description string `form:"description" binding:"required"` // 广告位描述
	ItemId      int    `form:"item_id"     binding:"required"` // 站点ID
	CateId      int    `form:"cate_id"     binding:"required"` // 栏目ID
	LocId       int    `form:"loc_id"      binding:"required"` // 广告页面位置
	Platform    int    `form:"platform"    binding:"required"` // 站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端
	Sort        int    `form:"sort"        binding:"required"` // 广告位排序
}
