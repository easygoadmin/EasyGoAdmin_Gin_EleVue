package model

import (
	"easygoadmin/library/db"
	"time"
)

type Example2 struct {
	Id         int       `xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `xorm:"default 'NULL' comment('演示名称') index VARCHAR(30)"`
	Status     int       `xorm:"default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Sort       int       `xorm:"default 0 comment('排序号') INT(11)"`
	CreateUser int       `xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"default 1 comment('有效标识') TINYINT(1)"`
}

func (Example2) TableName() string {
	return "sys_example2"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Example2) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Example2) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Example2) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Example2) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
