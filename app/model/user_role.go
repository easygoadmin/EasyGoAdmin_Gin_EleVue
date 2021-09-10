package model

import "easygoadmin/utils"

type UserRole struct {
	UserId int `json:"user_id" xorm:"not null default 0 comment('人员ID') index INT(10)"`
	RoleId int `json:"role_id" xorm:"not null default 0 comment('角色ID') INT(10)"`
}

// 根据条件查询单条数据
func (r *UserRole) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *UserRole) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

//批量删除
func (r *UserRole) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&UserRole{})
}
