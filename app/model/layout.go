package model

import (
	"easygoadmin/utils"
	"time"
)

type Layout struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	LayoutDescId int       `json:"layout_desc_id" xorm:"not null default 0 comment('布局描述ID') INT(10)"`
	Type         int       `json:"type" xorm:"not null default 0 comment('类型：1资讯文章') TINYINT(1)"`
	TypeId       int       `json:"type_id" xorm:"not null default 0 comment('对应的类型编号') INT(10)"`
	Image        string    `json:"image" xorm:"not null comment('图片路径') VARCHAR(150)"`
	Sort         int       `json:"sort" xorm:"not null default 125 comment('显示顺序') INT(11)"`
	CreateUser   int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime   time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser   int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(10)"`
	UpdateTime   time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
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
