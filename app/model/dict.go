package model

import (
	"easygoadmin/library/db"
	"time"
)

type Dict struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `xorm:"not null comment('字典名称') index VARCHAR(30)"`
	Code       string    `xorm:"not null comment('字典值') VARCHAR(50)"`
	Sort       int       `xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	Note       string    `xorm:"default 'NULL' comment('字典备注') VARCHAR(255)"`
	CreateUser int       `xorm:"default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

func (Dict) TableName() string {
	return "sys_dict"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Dict) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Dict) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Dict) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Dict) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
