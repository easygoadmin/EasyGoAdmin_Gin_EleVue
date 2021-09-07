package model

import (
	"easygoadmin/library/db"
	"time"
)

type City struct {
	Id         int64     `xorm:"pk autoincr comment('编号') BIGINT(11)"`
	Pid        int       `xorm:"not null default 0 comment('父级编号') INT(11)"`
	Level      int       `xorm:"not null default 0 comment('城市级别：1省 2市 3区') TINYINT(1)"`
	Name       string    `xorm:"not null comment('城市名称') index VARCHAR(50)"`
	Citycode   string    `xorm:"not null comment('城市编号（区号）') VARCHAR(10)"`
	PAdcode    string    `xorm:"default 'NULL' comment('父级地理编号') VARCHAR(10)"`
	Adcode     string    `xorm:"default 'NULL' comment('地理编号') VARCHAR(10)"`
	Lng        int       `xorm:"default NULL comment('城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7') INT(11)"`
	Lat        int       `xorm:"default NULL comment('城市坐标中心点纬度（* 1e6）') INT(11)"`
	Sort       int       `xorm:"not null default 125 comment('排序号') TINYINT(3)"`
	CreateUser int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime time.Time `xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int       `xorm:"not null default 1 comment('有效标记') TINYINT(1)"`
}

func (City) TableName() string {
	return "sys_city"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *City) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *City) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *City) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *City) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
