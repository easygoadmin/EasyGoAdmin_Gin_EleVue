package level

import (
	"easygoadmin/library/db"
	"time"
)

type Level struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('职级名称') index VARCHAR(30)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('显示顺序') INT(11)"`
	CreateUser int       `json:"createUser" xorm:"default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"createTime" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `json:"updateUser" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"updateTime" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

func TableName() string {
	return "sys_level"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Level) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(r)
}

// 插入数据
func (r *Level) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *Level) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Level) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.Id).Delete(r)
}
