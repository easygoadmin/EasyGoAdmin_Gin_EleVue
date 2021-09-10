package model

import (
	"easygoadmin/utils"
	"time"
)

type Item struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `json:"name" xorm:"not null comment('站点名称') index VARCHAR(15)"`
	Type       int       `json:"type" xorm:"not null default 1 comment('站点类型:1普通站点 2其他') TINYINT(3)"`
	Url        string    `json:"url" xorm:"not null comment('站点地址') VARCHAR(60)"`
	Image      string    `json:"image" xorm:"not null comment('站点图片') VARCHAR(100)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Note       string    `json:"note" xorm:"not null comment('站点备注') VARCHAR(255)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Item) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Item) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Item) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Item) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Item{})
}

//批量删除
func (r *Item) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Item{})
}
