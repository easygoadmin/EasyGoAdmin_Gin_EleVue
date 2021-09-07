package model

import (
	"easygoadmin/library/db"
	"time"
)

type ItemCate struct {
	Id         int       `xorm:"not null pk autoincr comment('ID') INT(11)"`
	Name       string    `xorm:"default 'NULL' comment('栏目名称') VARCHAR(30)"`
	Pid        int       `xorm:"default 0 comment('父级ID') index INT(11)"`
	ItemId     int       `xorm:"default 0 comment('站点ID') index INT(11)"`
	Pinyin     string    `xorm:"default 'NULL' comment('拼音(全)') VARCHAR(50)"`
	Code       string    `xorm:"default 'NULL' comment('拼音(简)') VARCHAR(10)"`
	IsCover    int       `xorm:"not null comment('是否有封面：1是 2否') TINYINT(1)"`
	Cover      string    `xorm:"default 'NULL' comment('封面') VARCHAR(50)"`
	Status     int       `xorm:"default 1 comment('状态：1启用 2停用') TINYINT(1)"`
	Note       string    `xorm:"default 'NULL' comment('备注') VARCHAR(200)"`
	Sort       int       `xorm:"default 125 comment('排序') INT(11)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `xorm:"not null comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (ItemCate) TableName() string {
	return "sys_item_cate"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *ItemCate) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *ItemCate) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *ItemCate) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *ItemCate) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
