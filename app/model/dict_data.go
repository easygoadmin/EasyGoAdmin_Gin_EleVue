package model

import (
	"easygoadmin/utils"
	"time"
)

type DictData struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('字典项名称') unique VARCHAR(50)"`
	Code       string    `json:"code" xorm:"not null comment('字典项值') VARCHAR(50)"`
	DictId     int       `json:"dict_id" xorm:"not null default 0 comment('字典类型ID') INT(11)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：1在用 2停用') TINYINT(1)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(300)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('显示顺序') SMALLINT(5)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标记') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *DictData) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *DictData) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *DictData) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *DictData) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&DictData{})
}

//批量删除
func (r *DictData) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&DictData{})
}
