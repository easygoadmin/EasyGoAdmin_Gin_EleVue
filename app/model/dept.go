package model

import (
	"easygoadmin/library/db"
	"time"
)

type Dept struct {
	Id         int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `xorm:"not null comment('部门名称') index VARCHAR(50)"`
	Code       string    `xorm:"default 'NULL' comment('部门编码') VARCHAR(150)"`
	Fullname   string    `xorm:"default 'NULL' comment('部门全称') VARCHAR(150)"`
	Type       int       `xorm:"not null default 0 comment('类型：1公司 2子公司 3部门 4小组') TINYINT(1)"`
	Pid        int       `xorm:"not null default 0 comment('上级ID') index INT(11)"`
	Sort       int       `xorm:"not null default 125 comment('排序') SMALLINT(5)"`
	Note       string    `xorm:"default 'NULL' comment('备注说明') VARCHAR(255)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

func (Dept) TableName() string {
	return "sys_dept"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Dept) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Dept) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Dept) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Dept) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
