package model

import (
	"easygoadmin/library/db"
	"time"
)

type Layout struct {
	Id           int       `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	LayoutDescId int       `xorm:"not null default 0 comment('布局描述ID') INT(10)"`
	Type         int       `xorm:"not null default 0 comment('类型：1资讯文章') TINYINT(1)"`
	TypeId       int       `xorm:"not null default 0 comment('对应的类型编号') INT(10)"`
	Image        string    `xorm:"not null comment('图片路径') VARCHAR(150)"`
	Sort         int       `xorm:"not null default 125 comment('显示顺序') INT(11)"`
	CreateUser   int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime   time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser   int       `xorm:"not null default 0 comment('更新人') INT(10)"`
	UpdateTime   time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (Layout) TableName() string {
	return "sys_layout"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Layout) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Layout) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Layout) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Layout) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
