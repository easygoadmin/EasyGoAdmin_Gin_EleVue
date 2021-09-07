package model

import (
	"easygoadmin/library/db"
	"time"
)

type Notice struct {
	Id         int       `xorm:"not null pk autoincr comment('通知ID') INT(11)"`
	Title      string    `xorm:"not null comment('通知标题') index VARCHAR(150)"`
	Content    string    `xorm:"not null comment('通知内容') TEXT"`
	Source     int       `xorm:"not null comment('来源：1内部通知 2外部通知') TINYINT(1)"`
	IsTop      int       `xorm:"not null default 2 comment('是否置顶：1是 2否') TINYINT(1)"`
	Browse     int       `xorm:"not null default 0 comment('阅读量') INT(10)"`
	Status     int       `xorm:"not null default 2 comment('状态：1已发布 2待发布') TINYINT(1)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

func (Notice) TableName() string {
	return "sys_notice"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Notice) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Notice) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Notice) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Notice) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
