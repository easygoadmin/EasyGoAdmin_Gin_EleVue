package model

import (
	"easygoadmin/library/db"
	"time"
)

type UmsMemberLevel struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `xorm:"not null comment('级别名称') index VARCHAR(30)"`
	Sort       int       `xorm:"not null default 125 comment('排序号') INT(11)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

func (UmsMemberLevel) TableName() string {
	return "sys_ums_member_level"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *UmsMemberLevel) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *UmsMemberLevel) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *UmsMemberLevel) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *UmsMemberLevel) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
