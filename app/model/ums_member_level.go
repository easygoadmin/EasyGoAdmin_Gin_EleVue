package model

import (
	"easygoadmin/utils"
	"time"
)

type UmsMemberLevel struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('级别名称') index VARCHAR(30)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('排序号') INT(11)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *UmsMemberLevel) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *UmsMemberLevel) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *UmsMemberLevel) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *UmsMemberLevel) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&UmsMemberLevel{})
}

//批量删除
func (r *UmsMemberLevel) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&UmsMemberLevel{})
}
