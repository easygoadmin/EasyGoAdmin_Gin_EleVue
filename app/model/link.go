package model

import (
	"easygoadmin/utils"
	"time"
)

type Link struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('唯一性标识') INT(10)"`
	Name       string    `json:"name" xorm:"default 'NULL' comment('友链名称') VARCHAR(50)"`
	Type       int       `json:"type" xorm:"not null default 1 comment('类型：1友情链接 2合作伙伴') TINYINT(1)"`
	Url        string    `json:"url" xorm:"default 'NULL' comment('友链地址') VARCHAR(150)"`
	ItemId     int       `json:"item_id" xorm:"not null default 0 comment('站点ID') INT(10)"`
	CateId     int       `json:"cate_id" xorm:"not null default 0 comment('栏目ID') INT(10)"`
	Platform   int       `json:"platform" xorm:"not null default 1 comment('平台：1PC站 2WAP站 3微信小程序 4APP应用') TINYINT(1)"`
	Form       int       `json:"form" xorm:"not null default 1 comment('友链形式：1文字链接 2图片链接') TINYINT(1)"`
	Image      string    `json:"image" xorm:"default 'NULL' comment('友链图片') VARCHAR(50)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(255)"`
	CreateUser int       `json:"create_user" xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') index TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Link) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Link) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Link) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Link) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Link{})
}

//批量删除
func (r *Link) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Link{})
}
