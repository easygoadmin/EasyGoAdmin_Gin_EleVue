package model

import (
	"easygoadmin/utils"
	"time"
)

type Notice struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('通知ID') INT(11)"`
	Title      string    `json:"title" xorm:"not null comment('通知标题') index VARCHAR(150)"`
	Content    string    `json:"content" xorm:"not null comment('通知内容') TEXT"`
	Source     int       `json:"source" xorm:"not null comment('来源：1内部通知 2外部通知') TINYINT(1)"`
	IsTop      int       `json:"isTop" xorm:"not null default 2 comment('是否置顶：1是 2否') TINYINT(1)"`
	Browse     int       `json:"browse" xorm:"not null default 0 comment('阅读量') INT(10)"`
	Status     int       `json:"status" xorm:"not null default 2 comment('状态：1已发布 2待发布') TINYINT(1)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Notice) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Notice) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Notice) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Notice) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Notice{})
}

//批量删除
func (r *Notice) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Notice{})
}
