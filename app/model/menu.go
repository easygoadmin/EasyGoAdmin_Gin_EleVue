package model

import (
	"easygoadmin/library/db"
	"time"
)

type Menu struct {
	Id         int       `xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `xorm:"not null comment('菜单名称') index VARCHAR(30)"`
	Icon       string    `xorm:"default 'NULL' comment('图标') VARCHAR(50)"`
	Url        string    `xorm:"default 'NULL' comment('URL地址') VARCHAR(150)"`
	Param      string    `xorm:"default 'NULL' comment('参数') VARCHAR(150)"`
	Pid        int       `xorm:"not null default 0 comment('上级ID') index INT(10)"`
	Type       int       `xorm:"not null default 0 comment('类型：1模块 2导航 3菜单 4节点') TINYINT(1)"`
	Permission string    `xorm:"default 'NULL' comment('权限标识') VARCHAR(150)"`
	Status     int       `xorm:"not null default 1 comment('是否显示：1显示 2不显示') TINYINT(1)"`
	Target     int       `xorm:"not null default 1 comment('打开方式：1内部打开 2外部打开') TINYINT(1)"`
	Note       string    `xorm:"default 'NULL' comment('菜单备注') VARCHAR(255)"`
	Sort       int       `xorm:"default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

func (Menu) TableName() string {
	return "sys_menu"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Menu) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Menu) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Menu) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Menu) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
