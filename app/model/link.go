package model

import (
	"easygoadmin/library/db"
	"time"
)

type Link struct {
	Id         int       `xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `xorm:"default 'NULL' comment('友链名称') VARCHAR(50)"`
	Type       int       `xorm:"not null default 1 comment('类型：1友情链接 2合作伙伴') TINYINT(1)"`
	Url        string    `xorm:"default 'NULL' comment('友链地址') VARCHAR(150)"`
	ItemId     int       `xorm:"not null default 0 comment('站点ID') INT(10)"`
	CateId     int       `xorm:"not null default 0 comment('栏目ID') INT(10)"`
	Platform   int       `xorm:"not null default 1 comment('平台：1PC站 2WAP站 3微信小程序 4APP应用') TINYINT(1)"`
	Form       int       `xorm:"not null default 1 comment('友链形式：1文字链接 2图片链接') TINYINT(1)"`
	Image      string    `xorm:"default 'NULL' comment('友链图片') VARCHAR(50)"`
	Status     int       `xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Sort       int       `xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') index TINYINT(1)"`
}

func (Link) TableName() string {
	return "sys_link"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Link) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Link) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Link) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Link) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
