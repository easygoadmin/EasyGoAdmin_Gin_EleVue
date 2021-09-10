package model

import (
	"easygoadmin/utils"
	"time"
)

type Menu struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	ParentId   int       `json:"parent_id" xorm:"not null default 0 comment('父级ID') index INT(11)"`
	Title      string    `json:"title" xorm:"not null comment('菜单标题') index VARCHAR(30)"`
	Icon       string    `json:"icon" xorm:"default 'NULL' comment('图标') VARCHAR(50)"`
	Path       string    `json:"path" xorm:"default 'NULL' comment('菜单路径') VARCHAR(150)"`
	Component  string    `json:"component" xorm:"default 'NULL' comment('菜单组件') VARCHAR(150)"`
	Target     string    `json:"target" xorm:"default 'NULL' comment('打开方式：0组件 1内链 2外链') VARCHAR(30)"`
	Permission string    `json:"permission" xorm:"default 'NULL' comment('权限标识') VARCHAR(150)"`
	Type       int       `json:"type" xorm:"not null default 0 comment('类型：0菜单 1节点') TINYINT(1)"`
	Method     string    `json:"method" xorm:"default 'NULL' comment('请求方式') VARCHAR(30)"`
	Status     int       `json:"status" xorm:"default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	Hide       int       `json:"hide" xorm:"default 1 comment('是否可见：1是 2否') TINYINT(1)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(255)"`
	Sort       int       `json:"sort" xorm:"default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Menu) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Menu) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Menu) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Menu) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Menu{})
}

//批量删除
func (r *Menu) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Menu{})
}
