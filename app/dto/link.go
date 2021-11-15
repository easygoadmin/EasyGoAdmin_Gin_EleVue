/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : link
 */
package dto

// 分页查询条件
type LinkPageReq struct {
	Name     string `form:"name"`     // 友链名称
	Type     int    `form:"type"`     // 友链类型
	Platform int    `form:"platform"` // 投放平台
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加友链
type LinkAddReq struct {
	Name     string `form:"name"        binding:"required"` // 友链名称
	Type     int    `form:"type"        binding:"required"` // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                            // 友链地址
	ItemId   int    `form:"item_id"`                        // 站点ID
	CateId   int    `form:"cate_id"`                        // 栏目ID
	Platform int    `form:"platform"    binding:"required"` // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form"        binding:"required"` // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                          // 友链图片
	Status   int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort     int    `form:"sort"        binding:"required"` // 显示顺序
	Note     string `form:"note"`                           // 备注
}

// 修改友链
type LinkUpdateReq struct {
	Id       int    `form:"id" v:"required#主键ID不能为空"`
	Name     string `form:"name"        binding:"required"` // 友链名称
	Type     int    `form:"type"        binding:"required"` // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                            // 友链地址
	ItemId   int    `form:"item_id"`                        // 站点ID
	CateId   int    `form:"cate_id"`                        // 栏目ID
	Platform int    `form:"platform"    binding:"required"` // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form"        binding:"required"` // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                          // 友链图片
	Status   int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort     int    `form:"sort"        binding:"required"` // 显示顺序
	Note     string `form:"note"`                           // 备注
}

// 设置状态
type LinkStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
