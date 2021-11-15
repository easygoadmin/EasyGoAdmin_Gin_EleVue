package model

import (
	"easygoadmin/utils"
	"time"
)

type City struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('编号') INT(11)"`
	Pid        int       `json:"pid" xorm:"not null default 0 comment('父级编号') INT(11)"`
	Level      int       `json:"level" xorm:"not null default 0 comment('城市级别：1省 2市 3区') TINYINT(1)"`
	Name       string    `json:"name" xorm:"not null comment('城市名称') index VARCHAR(50)"`
	Citycode   string    `json:"citycode" xorm:"not null comment('城市编号（区号）') VARCHAR(10)"`
	PAdcode    string    `json:"PAdcode" xorm:"default 'NULL' comment('父级地理编号') VARCHAR(10)"`
	Adcode     string    `json:"adcode" xorm:"default 'NULL' comment('地理编号') VARCHAR(10)"`
	Lng        int       `json:"lng" xorm:"default NULL comment('城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7') INT(11)"`
	Lat        int       `json:"lat" xorm:"default NULL comment('城市坐标中心点纬度（* 1e6）') INT(11)"`
	Sort       int       `json:"sort" xorm:"not null default 125 comment('排序号') TINYINT(3)"`
	CreateUser int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `json:"mark" xorm:"not null default 1 comment('有效标记') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *City) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *City) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *City) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *City) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&City{})
}

//批量删除
func (r *City) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&City{})
}
