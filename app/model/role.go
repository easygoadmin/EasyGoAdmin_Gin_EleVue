package model

import (
	"easygoadmin/library/db"
	"time"
)

type Role struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Name       string    `xorm:"not null comment('角色名称') index VARCHAR(30)"`
	Code       string    `xorm:"not null comment('角色标签') VARCHAR(100)"`
	Status     int       `xorm:"not null default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	Note       string    `xorm:"default 'NULL' comment('备注') VARCHAR(255)"`
	Sort       int       `xorm:"not null default 125 comment('排序') SMALLINT(5)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (Role) TableName() string {
	return "sys_role"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Role) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Role) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Role) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Role) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
