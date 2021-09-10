package model

import (
	"easygoadmin/utils"
	"time"
)

type Dept struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('部门名称') index VARCHAR(50)"`
	Code       string    `json:"code" xorm:"default 'NULL' comment('部门编码') VARCHAR(150)"`
	Fullname   string    `json:"fullname" xorm:"default 'NULL' comment('部门全称') VARCHAR(150)"`
	Type       int       `json:"type" xorm:"not null default 0 comment('类型：1公司 2子公司 3部门 4小组') TINYINT(1)"`
	Pid        int       `json:"pid" xorm:"not null default 0 comment('上级ID') index INT(11)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('排序') SMALLINT(5)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注说明') VARCHAR(255)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Dept) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Dept) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Dept) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Dept) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Dept{})
}

//批量删除
func (r *Dept) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Dept{})
}
