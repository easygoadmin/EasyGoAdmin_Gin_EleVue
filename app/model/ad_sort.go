package model

import (
	"easygoadmin/utils"
	"time"
)

type AdSort struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Description string    `json:"description" xorm:"not null comment('广告位描述') VARCHAR(255)"`
	ItemId      int       `json:"item_id" xorm:"not null default 0 comment('站点ID') INT(10)"`
	CateId      int       `json:"cate_id" xorm:"not null default 0 comment('栏目ID') SMALLINT(5)"`
	LocId       int       `json:"loc_id" xorm:"not null default 0 comment('广告页面位置') SMALLINT(5)"`
	Platform    int       `json:"platform" xorm:"not null default 1 comment('站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端') TINYINT(1)"`
	Sort        int       `json:"sort" xorm:"not null default 125 comment('广告位排序') SMALLINT(5)"`
	CreateUser  int       `json:"create_user" xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime  time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser  int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime  time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark        int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *AdSort) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *AdSort) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *AdSort) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *AdSort) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&AdSort{})
}

//批量删除
func (r *AdSort) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&AdSort{})
}
