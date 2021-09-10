package model

import (
	"easygoadmin/utils"
	"time"
)

type OperLog struct {
	Id           int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(20)"`
	Model        string    `json:"model" xorm:"not null comment('操作模块') VARCHAR(150)"`
	OperType     int       `json:"oper_type" xorm:"not null default 0 comment('操作类型：0其它 1新增 2修改 3删除 4查询 5设置状态 6导入 7导出 8设置权限 9设置密码') INT(2)"`
	OperMethod   string    `json:"oper_method" xorm:"default 'NULL' comment('操作方法') VARCHAR(30)"`
	Username     string    `json:"username" xorm:"default 'NULL' comment('操作账号') VARCHAR(255)"`
	OperName     string    `json:"oper_name" xorm:"default 'NULL' comment('操作用户') VARCHAR(50)"`
	OperUrl      string    `json:"oper_url" xorm:"default 'NULL' comment('请求URL') VARCHAR(255)"`
	OperIp       string    `json:"oper_ip" xorm:"default '''' comment('主机地址') VARCHAR(50)"`
	OperLocation string    `json:"oper_location" xorm:"default '''' comment('操作地点') VARCHAR(255)"`
	RequestParam string    `json:"request_param" xorm:"default '''' comment('请求参数') VARCHAR(2000)"`
	Result       string    `json:"result" xorm:"default '''' comment('返回参数') VARCHAR(2000)"`
	Status       int       `json:"status" xorm:"default 0 comment('日志状态：0正常日志 1错误日志') TINYINT(1)"`
	UserAgent    string    `json:"user_agent" xorm:"default 'NULL' comment('代理信息') TEXT"`
	Note         string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(2000)"`
	CreateUser   int       `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime   time.Time `json:"create_time" xorm:"default 'NULL' comment('操作时间') DATETIME"`
	UpdateUser   int       `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime   time.Time `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int       `json:"mark" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *OperLog) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *OperLog) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *OperLog) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *OperLog) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&OperLog{})
}

//批量删除
func (r *OperLog) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&OperLog{})
}
