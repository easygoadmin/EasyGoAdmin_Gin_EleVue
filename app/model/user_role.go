package model

import (
	"easygoadmin/library/db"
)

type UserRole struct {
	UserId int `xorm:"not null default 0 comment('人员ID') index INT(10)"`
	RoleId int `xorm:"not null default 0 comment('角色ID') INT(10)"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *UserRole) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *UserRole) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}
