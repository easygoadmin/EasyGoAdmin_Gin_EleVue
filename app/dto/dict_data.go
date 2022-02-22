// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
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
 * 字典数据Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : dict_data
 */
package dto

// 字典项列表查询条件
type DictDataPageReq struct {
	DictId int    `form:"dictId"` // 字典ID
	Name   string `form:"name"`   // 字典项名称
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加字典项
type DictDataAddReq struct {
	Name   string `form:"name,unique" binding:"required"` // 字典项名称
	Code   string `form:"code"        binding:"required"` // 字典项值
	DictId int    `form:"dict_id"     binding:"required"` // 字典类型ID
	Note   string `form:"note"`                           // 备注
	Sort   int    `form:"sort"`                           // 显示顺序
}

// 更新字典项
type DictDataUpdateReq struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name,unique" binding:"required"` // 字典项名称
	Code   string `form:"code"        binding:"required"` // 字典项值
	DictId int    `form:"dict_id"     binding:"required"` // 字典类型ID
	Note   string `form:"note"`                           // 备注
	Sort   int    `form:"sort"`                           // 显示顺序
}
