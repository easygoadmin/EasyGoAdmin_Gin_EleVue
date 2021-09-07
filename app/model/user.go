package model

import (
	"easygoadmin/library/db"
	"time"
)

type User struct {
	Id           int       `xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Realname     string    `xorm:"default 'NULL' comment('真实姓名') index VARCHAR(150)"`
	Nickname     string    `xorm:"default 'NULL' comment('昵称') VARCHAR(150)"`
	Gender       int       `xorm:"default 3 comment('性别:1男 2女 3保密') TINYINT(1)"`
	Avatar       string    `xorm:"default 'NULL' comment('头像') VARCHAR(150)"`
	Mobile       string    `xorm:"default 'NULL' comment('手机号码') CHAR(11)"`
	Email        string    `xorm:"default 'NULL' comment('邮箱地址') VARCHAR(30)"`
	Birthday     time.Time `xorm:"default 'NULL' comment('出生日期') DATE"`
	DeptId       int       `xorm:"default 0 comment('部门ID') INT(11)"`
	LevelId      int       `xorm:"default 0 comment('职级ID') INT(11)"`
	PositionId   int       `xorm:"default 0 comment('岗位ID') SMALLINT(3)"`
	ProvinceCode string    `xorm:"default 'NULL' comment('省份编号') VARCHAR(50)"`
	CityCode     string    `xorm:"default 'NULL' comment('市区编号') VARCHAR(50)"`
	DistrictCode string    `xorm:"default 'NULL' comment('区县编号') VARCHAR(50)"`
	Address      string    `xorm:"default 'NULL' comment('详细地址') VARCHAR(255)"`
	CityName     string    `xorm:"default 'NULL' comment('所属城市') VARCHAR(150)"`
	Username     string    `xorm:"default 'NULL' comment('登录用户名') VARCHAR(50)"`
	Password     string    `xorm:"default 'NULL' comment('登录密码') VARCHAR(150)"`
	Salt         string    `xorm:"default 'NULL' comment('盐加密') VARCHAR(30)"`
	Intro        string    `xorm:"default 'NULL' comment('个人简介') VARCHAR(500)"`
	Status       int       `xorm:"default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	Note         string    `xorm:"default 'NULL' comment('备注') VARCHAR(500)"`
	Sort         int       `xorm:"default 125 comment('排序号') INT(11)"`
	LoginNum     int       `xorm:"default 0 comment('登录次数') INT(11)"`
	LoginIp      string    `xorm:"default 'NULL' comment('最近登录IP') VARCHAR(20)"`
	LoginTime    time.Time `xorm:"default 'NULL' comment('最近登录时间') DATETIME"`
	CreateUser   int       `xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime   time.Time `xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser   int       `xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime   time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int       `xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

func (User) TableName() string {
	return "sys_user"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *User) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *User) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *User) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *User) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
