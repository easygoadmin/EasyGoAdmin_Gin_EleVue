// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨
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
 * 友链Dto
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
	Sort     int    `form:"sort"`                           // 显示顺序
	Note     string `form:"note"`                           // 备注
}

// 修改友链
type LinkUpdateReq struct {
	Id       int    `form:"id" binding:"required"`
	Name     string `form:"name"        binding:"required"` // 友链名称
	Type     int    `form:"type"        binding:"required"` // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                            // 友链地址
	ItemId   int    `form:"item_id"`                        // 站点ID
	CateId   int    `form:"cate_id"`                        // 栏目ID
	Platform int    `form:"platform"    binding:"required"` // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form"        binding:"required"` // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                          // 友链图片
	Status   int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Sort     int    `form:"sort"`                           // 显示顺序
	Note     string `form:"note"`                           // 备注
}

// 设置状态
type LinkStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
