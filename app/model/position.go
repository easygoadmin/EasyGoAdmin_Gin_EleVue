package model

import (
	"easygoadmin/library/db"
	"time"
)

type Position struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `xorm:"default 'NULL' comment('岗位名称') index VARCHAR(30)"`
	Status     int       `xorm:"default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Sort       int       `xorm:"default 125 comment('显示顺序') INT(11)"`
	CreateUser int       `xorm:"default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"default 1 comment('有效标识') TINYINT(1)"`
}

func (Position) TableName() string {
	return "sys_position"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Position) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Position) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Position) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Position) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
