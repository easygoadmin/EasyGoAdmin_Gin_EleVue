package model

import (
	"easygoadmin/library/db"
	"time"
)

type ConfigData struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Title      string    `xorm:"not null comment('配置标题') index VARCHAR(50)"`
	Code       string    `xorm:"not null comment('配置编码') index VARCHAR(100)"`
	Value      string    `xorm:"default 'NULL' comment('配置值') TEXT"`
	Options    string    `xorm:"default 'NULL' comment('配置项') VARCHAR(255)"`
	ConfigId   int       `xorm:"not null default 0 comment('配置ID') INT(11)"`
	Type       string    `xorm:"not null comment('配置类型') VARCHAR(16)"`
	Status     int       `xorm:"not null default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Sort       int       `xorm:"not null default 0 comment('排序') SMALLINT(3)"`
	Note       string    `xorm:"default 'NULL' comment('配置说明') VARCHAR(500)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识：1正常 0删除') TINYINT(1)"`
}

func (ConfigData) TableName() string {
	return "sys_config_data"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *ConfigData) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *ConfigData) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *ConfigData) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *ConfigData) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
