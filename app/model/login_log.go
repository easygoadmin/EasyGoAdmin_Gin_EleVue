package model

import (
	"easygoadmin/utils"
	"time"
)

type LoginLog struct {
	Id           int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(20)"`
	Username     string    `json:"username" xorm:"default 'NULL' comment('操作账号') VARCHAR(255)"`
	Method       string    `json:"method" xorm:"default 'NULL' comment('操作方法') VARCHAR(30)"`
	OperUrl      string    `json:"oper_url" xorm:"default 'NULL' comment('请求URL') VARCHAR(255)"`
	OperIp       string    `json:"oper_ip" xorm:"default '''' comment('主机地址') VARCHAR(50)"`
	OperLocation string    `json:"oper_location" xorm:"default '''' comment('操作地点') VARCHAR(255)"`
	Os           string    `json:"os" xorm:"default 'NULL' comment('操作系统') VARCHAR(255)"`
	RequestParam string    `json:"request_param" xorm:"default '''' comment('请求参数') VARCHAR(2000)"`
	Browser      string    `json:"browser" xorm:"default 'NULL' comment('浏览器') VARCHAR(255)"`
	Result       string    `json:"result" xorm:"default '''' comment('返回参数') VARCHAR(2000)"`
	Status       int       `json:"status" xorm:"default 1 comment('操作状态：0操作成功 1操作失败') TINYINT(1)"`
	Type         int       `json:"type" xorm:"default 0 comment('操作类型：1登录成功 2登录失败 3注销成功 2注销失败') TINYINT(1)"`
	UserAgent    string    `json:"user_agent" xorm:"default 'NULL' comment('代理信息') TEXT"`
	Note         string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(2000)"`
	CreateUser   int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime   time.Time `json:"create_time" xorm:"default 'NULL' comment('操作时间') DATETIME"`
	UpdateUser   int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime   time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *LoginLog) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *LoginLog) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *LoginLog) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *LoginLog) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&LoginLog{})
}

//批量删除
func (r *LoginLog) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&LoginLog{})
}
