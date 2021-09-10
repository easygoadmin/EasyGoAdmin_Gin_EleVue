package model

import (
	"easygoadmin/utils"
	"time"
)

type Ad struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Title       string    `json:"title" xorm:"not null comment('广告标题') index VARCHAR(100)"`
	AdSortId    int       `json:"ad_sort_id" xorm:"not null default 0 comment('广告位ID') index INT(11)"`
	Cover       string    `json:"cover" xorm:"default 'NULL' comment('广告图片') VARCHAR(255)"`
	Type        int       `json:"type" xorm:"not null default 0 comment('广告格式：1图片 2文字 3视频 4推荐') TINYINT(1)"`
	Description string    `json:"description" xorm:"default 'NULL' comment('广告描述') VARCHAR(150)"`
	Content     string    `json:"content" xorm:"default 'NULL' comment('广告内容') TEXT"`
	Url         string    `json:"url" xorm:"default 'NULL' comment('广告链接') TEXT"`
	Width       int       `json:"width" xorm:"not null default 0 comment('广告宽度') INT(10)"`
	Height      int       `json:"height" xorm:"not null default 0 comment('广告高度') INT(10)"`
	StartTime   time.Time `json:"start_time" xorm:"default 'NULL' comment('开始时间') DATETIME"`
	EndTime     time.Time `json:"end_time" xorm:"default 'NULL' comment('结束时间') DATETIME"`
	ViewNum     int       `json:"view_num" xorm:"not null default 0 comment('点击率') INT(10)"`
	Status      int       `json:"status" xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Sort        int       `json:"sort" xorm:"not null default 125 comment('排序') SMALLINT(5)"`
	CreateUser  int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null comment('添加时间') DATETIME"`
	UpdateUser  int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime  time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark        int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Ad) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Ad) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Ad) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Ad) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Ad{})
}

//批量删除
func (r *Ad) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Ad{})
}
