package model

import (
	"easygoadmin/utils"
	"time"
)

type Example struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('唯一性标识') INT(11)"`
	Name       string    `json:"name" xorm:"default 'NULL' comment('测试名称') index VARCHAR(30)"`
	Avatar     string    `json:"avatar" xorm:"default 'NULL' comment('头像') VARCHAR(255)"`
	Content    string    `json:"content" xorm:"default 'NULL' comment('内容') VARCHAR(255)"`
	Status     int       `json:"status" xorm:"default 1 comment('状态：1正常 2停用') TINYINT(1)"`
	Type       int       `json:"type" xorm:"default 1 comment('类型：1京东 2淘宝 3拼多多 4唯品会') INT(11)"`
	IsVip      int       `json:"is_vip" xorm:"default 2 comment('是否VIP：1是 2否') TINYINT(1)"`
	Sort       int       `json:"sort" xorm:"default 0 comment('排序号') INT(11)"`
	CreateUser int       `json:"create_user" xorm:"default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Example) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Example) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Example) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Example) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Example{})
}

//批量删除
func (r *Example) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Example{})
}
