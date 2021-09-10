package model

import (
	"easygoadmin/utils"
	"time"
)

type Position struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `json:"name" xorm:"default 'NULL' comment('岗位名称') index VARCHAR(30)"`
	Status     int       `json:"status" xorm:"default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Sort       int       `json:"sort" xorm:"default 125 comment('显示顺序') INT(11)"`
	CreateUser int       `json:"create_user" xorm:"default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Position) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Position) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Position) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Position) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Position{})
}

//批量删除
func (r *Position) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Position{})
}
