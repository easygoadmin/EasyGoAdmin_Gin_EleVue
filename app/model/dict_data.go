package model

import (
	"easygoadmin/library/db"
	"time"
)

type DictData struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `xorm:"not null comment('字典项名称') unique VARCHAR(50)"`
	Code       string    `xorm:"not null comment('字典项值') VARCHAR(50)"`
	DictId     int       `xorm:"not null default 0 comment('字典类型ID') INT(11)"`
	Status     int       `xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Note       string    `xorm:"default 'NULL' comment('备注') VARCHAR(300)"`
	Sort       int       `xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标记') TINYINT(1)"`
}

func (DictData) TableName() string {
	return "sys_dict_data"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *DictData) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *DictData) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *DictData) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *DictData) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
