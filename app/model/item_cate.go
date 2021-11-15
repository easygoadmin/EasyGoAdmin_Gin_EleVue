package model

import (
	"easygoadmin/utils"
	"time"
)

type ItemCate struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('ID') INT(11)"`
	Name       string    `json:"name" xorm:"default 'NULL' comment('栏目名称') VARCHAR(30)"`
	Pid        int       `json:"pid" xorm:"default 0 comment('父级ID') index INT(11)"`
	ItemId     int       `json:"itemId" xorm:"default 0 comment('站点ID') index INT(11)"`
	Pinyin     string    `json:"pinyin" xorm:"default 'NULL' comment('拼音(全)') VARCHAR(50)"`
	Code       string    `json:"code" xorm:"default 'NULL' comment('拼音(简)') VARCHAR(10)"`
	IsCover    int       `json:"isCover" xorm:"not null comment('是否有封面：1是 2否') TINYINT(1)"`
	Cover      string    `json:"cover" xorm:"default 'NULL' comment('封面') VARCHAR(50)"`
	Status     int       `json:"status" xorm:"default 1 comment('状态：1启用 2停用') TINYINT(1)"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(200)"`
	Sort       int       `json:"sort" xorm:"default 125 comment('排序') INT(11)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(10)"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *ItemCate) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *ItemCate) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *ItemCate) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *ItemCate) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&ItemCate{})
}

//批量删除
func (r *ItemCate) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&ItemCate{})
}
