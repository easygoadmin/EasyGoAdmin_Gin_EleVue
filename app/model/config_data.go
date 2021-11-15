package model

import (
	"easygoadmin/utils"
	"time"
)

type ConfigData struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Title      string    `json:"title" xorm:"not null comment('配置标题') index VARCHAR(50)"`
	Code       string    `json:"code" xorm:"not null comment('配置编码') index VARCHAR(100)"`
	Value      string    `json:"value" xorm:"default 'NULL' comment('配置值') TEXT"`
	Options    string    `json:"options" xorm:"default 'NULL' comment('配置项') VARCHAR(255)"`
	ConfigId   int       `json:"configId" xorm:"not null default 0 comment('配置ID') INT(11)"`
	Type       string    `json:"type" xorm:"not null comment('配置类型') VARCHAR(16)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Sort       int       `json:"sort" xorm:"not null default 0 comment('排序') SMALLINT(3)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('配置说明') VARCHAR(500)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识：1正常 0删除') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *ConfigData) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *ConfigData) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *ConfigData) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *ConfigData) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&ConfigData{})
}

//批量删除
func (r *ConfigData) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&ConfigData{})
}
