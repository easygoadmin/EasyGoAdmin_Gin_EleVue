package model

import (
	"easygoadmin/library/db"
	"time"
)

type Example struct {
	Id         int       `xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `xorm:"default 'NULL' comment('测试名称') index VARCHAR(30)"`
	Avatar     string    `xorm:"default 'NULL' comment('头像') VARCHAR(255)"`
	Content    string    `xorm:"default 'NULL' comment('内容') VARCHAR(255)"`
	Status     int       `xorm:"default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Type       int       `xorm:"default 1 comment('类型：1京东 2淘宝 3拼多多 4唯品会') INT(10)"`
	IsVip      int       `xorm:"default 2 comment('是否VIP：1是 2否') TINYINT(1)"`
	Sort       int       `xorm:"default 0 comment('排序号') INT(11)"`
	CreateUser int       `xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"default 1 comment('有效标识') TINYINT(1)"`
}

func (Example) TableName() string {
	return "sys_example"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *Example) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *Example) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *Example) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *Example) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
