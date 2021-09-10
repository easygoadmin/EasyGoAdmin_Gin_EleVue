package model

import (
	"easygoadmin/utils"
	"time"
)

type LayoutDesc struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	LocDesc    string    `json:"loc_desc" xorm:"not null comment('页面位置描述') VARCHAR(255)"`
	LocId      int       `json:"loc_id" xorm:"not null default 0 comment('位置编号') INT(11)"`
	ItemId     int       `json:"item_id" xorm:"not null default 0 comment('站点ID') INT(10)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('排序号') INT(11)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *LayoutDesc) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *LayoutDesc) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *LayoutDesc) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *LayoutDesc) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&LayoutDesc{})
}

//批量删除
func (r *LayoutDesc) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&LayoutDesc{})
}
