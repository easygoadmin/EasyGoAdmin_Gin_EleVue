package model

import (
	"easygoadmin/library/db"
	"time"
)

type AdSort struct {
	Id          int       `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Description string    `xorm:"not null comment('广告位描述') VARCHAR(255)"`
	ItemId      int       `xorm:"not null default 0 comment('站点ID') INT(10)"`
	CateId      int       `xorm:"not null default 0 comment('栏目ID') SMALLINT(5)"`
	LocId       int       `xorm:"not null default 0 comment('广告页面位置') SMALLINT(5)"`
	Platform    int       `xorm:"not null default 1 comment('站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端') TINYINT(1)"`
	Sort        int       `xorm:"not null default 125 comment('广告位排序') SMALLINT(5)"`
	CreateUser  int       `xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime  time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser  int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime  time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark        int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (AdSort) TableName() string {
	return "sys_ad_sort"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *AdSort) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *AdSort) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *AdSort) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *AdSort) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
