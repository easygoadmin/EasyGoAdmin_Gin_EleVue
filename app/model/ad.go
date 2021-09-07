package model

import (
	"easygoadmin/library/db"
	"time"
)

type Ad struct {
	Id          int       `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Title       string    `xorm:"not null comment('广告标题') index VARCHAR(100)"`
	AdSortId    int       `xorm:"not null default 0 comment('广告位ID') index INT(11)"`
	Cover       string    `xorm:"default 'NULL' comment('广告图片') VARCHAR(255)"`
	Type        int       `xorm:"not null default 0 comment('广告格式：1图片 2文字 3视频 4推荐') TINYINT(1)"`
	Description string    `xorm:"default 'NULL' comment('广告描述') VARCHAR(150)"`
	Content     string    `xorm:"default 'NULL' comment('广告内容') TEXT"`
	Url         string    `xorm:"default 'NULL' comment('广告链接') TEXT"`
	Width       int       `xorm:"not null default 0 comment('广告宽度') INT(10)"`
	Height      int       `xorm:"not null default 0 comment('广告高度') INT(10)"`
	StartTime   time.Time `xorm:"default 'NULL' comment('开始时间') DATETIME"`
	EndTime     time.Time `xorm:"default 'NULL' comment('结束时间') DATETIME"`
	ViewNum     int       `xorm:"not null default 0 comment('点击率') INT(10)"`
	Status      int       `xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Sort        int       `xorm:"not null default 125 comment('排序') SMALLINT(5)"`
	CreateUser  int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime  time.Time `xorm:"not null comment('添加时间') DATETIME"`
	UpdateUser  int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime  time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark        int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (Ad) TableName() string {
	return "sys_ad"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Ad) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Ad) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Ad) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Ad) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
