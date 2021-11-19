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
 * 站点Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : item
 */
package dto

// 分页查询条件
type ItemPageReq struct {
	Name  string `form:"name"`  // 站点名称
	Type  int    `form:"type"`  // 站点类型
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加站点
type ItemAddReq struct {
	Name   string `form:"name"        binding:"required"` // 站点名称
	Type   int    `form:"type"        binding:"required"` // 站点类型:1普通站点 2其他
	Url    string `form:"url"         binding:"required"` // 站点地址
	Image  string `form:"image"`                          // 站点图片
	Status int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Note   string `form:"note"`                           // 站点备注
	Sort   int    `form:"sort"        binding:"required"` // 显示顺序
}

// 更新站点
type ItemUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name"        binding:"required"` // 站点名称
	Type   int    `form:"type"        binding:"required"` // 站点类型:1普通站点 2其他
	Url    string `form:"url"         binding:"required"` // 站点地址
	Image  string `form:"image"`                          // 站点图片
	Status int    `form:"status"      binding:"required"` // 状态：1在用 2停用
	Note   string `form:"note"`                           // 站点备注
	Sort   int    `form:"sort"        binding:"required"` // 显示顺序
}

// 设置状态
type ItemStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
