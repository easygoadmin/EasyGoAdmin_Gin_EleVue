package model

import (
	"easygoadmin/library/db"
)

type RoleMenu struct {
	RoleId int `xorm:"not null default 0 comment('角色ID') SMALLINT(5)"`
	MenuId int `xorm:"not null default 0 comment('菜单ID') index SMALLINT(5)"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *RoleMenu) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *RoleMenu) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}
