package model

import (
	"easygoadmin/utils"
	"time"
)

type Config struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('分组名称') index VARCHAR(30)"`
	Sort       int       `json:"sort" xorm:"not null comment('排序') TINYINT(2)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Config) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Config) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Config) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Config) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Config{})
}

//批量删除
func (r *Config) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Config{})
}
