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

package model

import (
	"easygoadmin/utils"
)

type Layout struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	LayoutDescId int    `json:"layout_desc_id" xorm:"not null default 0 comment('布局描述ID') INT(10)"`
	Type         int    `json:"type" xorm:"not null default 0 comment('类型：1资讯文章') TINYINT(1)"`
	TypeId       int    `json:"type_id" xorm:"not null default 0 comment('对应的类型编号') INT(10)"`
	Image        string `json:"image" xorm:"not null comment('图片路径') VARCHAR(150)"`
	Sort         int    `json:"sort" xorm:"not null default 125 comment('显示顺序') INT(11)"`
	CreateUser   int    `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime   int64  `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser   int    `json:"update_user" xorm:"not null default 0 comment('更新人') INT(10)"`
	UpdateTime   int64  `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int    `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Layout) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Layout) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Layout) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Layout) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Layout{})
}

//批量删除
func (r *Layout) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Layout{})
}
