package model

import (
	"easygoadmin/library/db"
	"time"
)

type Item struct {
	Id         int       `xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `xorm:"not null comment('站点名称') index VARCHAR(15)"`
	Type       int       `xorm:"not null default 1 comment('站点类型:1普通站点 2其他') TINYINT(3)"`
	Url        string    `xorm:"not null comment('站点地址') VARCHAR(60)"`
	Image      string    `xorm:"not null comment('站点图片') VARCHAR(100)"`
	Status     int       `xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Note       string    `xorm:"not null comment('站点备注') VARCHAR(255)"`
	Sort       int       `xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"not null comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (Item) TableName() string {
	return "sys_item"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Item) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Item) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Item) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Item) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
