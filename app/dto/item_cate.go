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
 * 栏目Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item_cate
 */
package dto

// 栏目查询条件
type ItemCateQueryReq struct {
	Name string `form:"name"` // 栏目名称
}

// 添加站点
type ItemCateAddReq struct {
	Name    string `form:"name"        binding:"required"` // 栏目名称
	Pid     int    `form:"pid"`                            // 父级ID
	ItemId  int    `form:"item_id"     binding:"required"` // 站点ID
	Pinyin  string `form:"pinyin"      binding:"required"` // 拼音(全)
	Code    string `form:"code"        binding:"required"` // 拼音(简)
	IsCover int    `form:"is_cover"    binding:"required"` // 是否有封面：1是 2否
	Cover   string `form:"cover"`                          // 封面
	Status  int    `form:"status"      binding:"required"` // 状态：1启用 2停用
	Note    string `form:"note"`                           // 备注
	Sort    int    `form:"sort"        binding:"required"` // 排序
}

// 修改
type ItemCateUpdateReq struct {
	Id      int    `form:"id" binding:"required"`
	Name    string `form:"name"        binding:"required"` // 栏目名称
	Pid     int    `form:"pid"`                            // 父级ID
	ItemId  int    `form:"item_id"     binding:"required"` // 站点ID
	Pinyin  string `form:"pinyin"      binding:"required"` // 拼音(全)
	Code    string `form:"code"        binding:"required"` // 拼音(简)
	IsCover int    `form:"is_cover"    binding:"required"` // 是否有封面：1是 2否
	Cover   string `form:"cover"`                          // 封面
	Status  int    `form:"status"      binding:"required"` // 状态：1启用 2停用
	Note    string `form:"note"`                           // 备注
	Sort    int    `form:"sort"        binding:"required"` // 排序
}
