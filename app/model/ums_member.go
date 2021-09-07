package model

import (
	"easygoadmin/library/db"
	"time"
)

type UmsMember struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Openid       string    `xorm:"default 'NULL' comment('用户唯一标识') VARCHAR(50)"`
	Username     string    `xorm:"not null comment('用户名') index unique VARCHAR(30)"`
	Password     string    `xorm:"default 'NULL' comment('密码') CHAR(32)"`
	MemberLevel  int       `xorm:"not null default 0 comment('会员等级') SMALLINT(3)"`
	Realname     string    `xorm:"default 'NULL' comment('真实姓名') VARCHAR(50)"`
	Nickname     string    `xorm:"default 'NULL' comment('用户昵称') VARCHAR(50)"`
	Gender       int       `xorm:"not null default 3 comment('性别（1男 2女 3未知）') TINYINT(1)"`
	Avatar       string    `xorm:"default '''' comment('用户头像') VARCHAR(180)"`
	Birthday     time.Time `xorm:"default 'NULL' comment('出生日期') DATE"`
	ProvinceCode string    `xorm:"default 'NULL' comment('户籍省份编号') VARCHAR(30)"`
	CityCode     string    `xorm:"default 'NULL' comment('户籍城市编号') VARCHAR(30)"`
	DistrictCode string    `xorm:"default 'NULL' comment('户籍区/县编号') VARCHAR(30)"`
	Address      string    `xorm:"default 'NULL' comment('详细地址') VARCHAR(255)"`
	Intro        string    `xorm:"default 'NULL' comment('个人简介') TEXT"`
	Signature    string    `xorm:"default 'NULL' comment('个性签名') VARCHAR(30)"`
	Device       int       `xorm:"not null default 0 comment('设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加') TINYINT(1)"`
	DeviceCode   string    `xorm:"default 'NULL' comment('推送的别名') VARCHAR(40)"`
	PushAlias    string    `xorm:"default '''' comment('推送的别名') VARCHAR(40)"`
	Source       int       `xorm:"not null default 1 comment('来源：1、APP注册；2、后台添加；') TINYINT(1)"`
	Status       int       `xorm:"not null default 1 comment('是否启用 1、启用  2、停用') TINYINT(1)"`
	AppVersion   string    `xorm:"default '''' comment('客户端版本号') VARCHAR(30)"`
	Code         string    `xorm:"default 'NULL' comment('我的推广码') VARCHAR(10)"`
	LoginIp      string    `xorm:"default 'NULL' comment('最近登录IP') VARCHAR(30)"`
	LoginTime    time.Time `xorm:"default 'NULL' comment('登录时间') DATETIME"`
	LoginRegion  string    `xorm:"default 'NULL' comment('上次登录地点') VARCHAR(20)"`
	LoginCount   int       `xorm:"not null default 0 comment('登录总次数') INT(10)"`
	CreateUser   int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime   time.Time `xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser   int       `xorm:"not null default 0 comment('修改人') INT(11)"`
	UpdateTime   time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int       `xorm:"not null default 1 comment('有效标识：1正常 0删除') TINYINT(1)"`
}

func (UmsMember) TableName() string {
	return "sys_ums_member"
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *UmsMember) FindOne() (bool, error) {
	return db.Instance().Engine().Table(r.TableName()).Get(r)
}

// 插入数据
func (r *UmsMember) Insert() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *UmsMember) Update() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Update(r)
}

// 删除
func (r *UmsMember) Delete() (int64, error) {
	return db.Instance().Engine().Table(r.TableName()).ID(r.Id).Delete(r)
}
