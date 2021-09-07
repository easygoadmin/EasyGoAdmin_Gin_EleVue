package model

import (
	"easygoadmin/library/db"
	"time"
)

type LayoutDesc struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	LocDesc    string    `xorm:"not null comment('页面位置描述') VARCHAR(255)"`
	LocId      int       `xorm:"not null default 0 comment('位置编号') INT(11)"`
	ItemId     int       `xorm:"not null default 0 comment('站点ID') INT(10)"`
	Sort       int       `xorm:"not null default 125 comment('排序号') INT(11)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"not null default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (LayoutDesc) TableName() string {
	return "sys_layout_desc"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *LayoutDesc) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *LayoutDesc) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *LayoutDesc) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *LayoutDesc) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
