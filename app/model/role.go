package model

import (
	"easygoadmin/utils"
	"time"
)

type Role struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Name       string    `json:"name" xorm:"not null comment('角色名称') index VARCHAR(30)"`
	Code       string    `json:"code" xorm:"not null comment('角色标签') VARCHAR(100)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(255)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('排序') SMALLINT(5)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Role) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Role) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Role) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Role) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Role{})
}

//批量删除
func (r *Role) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Role{})
}
