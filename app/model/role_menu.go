package model

import "easygoadmin/utils"

type RoleMenu struct {
	RoleId int `json:"role_id" xorm:"not null default 0 comment('角色ID') SMALLINT(5)"`
	MenuId int `json:"menu_id" xorm:"not null default 0 comment('菜单ID') index SMALLINT(5)"`
}

// 根据条件查询单条数据
func (r *RoleMenu) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *RoleMenu) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

//批量删除
func (r *RoleMenu) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&RoleMenu{})
}
