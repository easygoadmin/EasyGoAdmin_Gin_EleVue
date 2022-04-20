// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package model

import (
	"easygoadmin/utils"
)

type Member struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Openid       string `json:"openid" xorm:"default 'NULL' comment('用户唯一标识') VARCHAR(50)"`
	Username     string `json:"username" xorm:"not null comment('用户名') index unique VARCHAR(30)"`
	Password     string `json:"password" xorm:"default 'NULL' comment('密码') CHAR(32)"`
	MemberLevel  int    `json:"memberLevel" xorm:"not null default 0 comment('会员等级') SMALLINT(3)"`
	Realname     string `json:"realname" xorm:"default 'NULL' comment('真实姓名') VARCHAR(50)"`
	Nickname     string `json:"nickname" xorm:"default 'NULL' comment('用户昵称') VARCHAR(50)"`
	Gender       int    `json:"gender" xorm:"not null default 3 comment('性别（1男 2女 3未知）') TINYINT(1)"`
	Avatar       string `json:"avatar" xorm:"default '''' comment('用户头像') VARCHAR(180)"`
	Birthday     int64  `json:"birthday" xorm:"default 'NULL' comment('出生日期') DATE"`
	ProvinceCode string `json:"province_code" xorm:"default 'NULL' comment('户籍省份编号') VARCHAR(30)"`
	CityCode     string `json:"city_code" xorm:"default 'NULL' comment('户籍城市编号') VARCHAR(30)"`
	DistrictCode string `json:"district_code" xorm:"default 'NULL' comment('户籍区/县编号') VARCHAR(30)"`
	Address      string `json:"address" xorm:"default 'NULL' comment('详细地址') VARCHAR(255)"`
	Intro        string `json:"intro" xorm:"default 'NULL' comment('个人简介') TEXT"`
	Signature    string `json:"signature" xorm:"default 'NULL' comment('个性签名') VARCHAR(30)"`
	Device       int    `json:"device" xorm:"not null default 0 comment('设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加') TINYINT(1)"`
	DeviceCode   string `json:"device_code" xorm:"default 'NULL' comment('推送的别名') VARCHAR(40)"`
	PushAlias    string `json:"push_alias" xorm:"default '''' comment('推送的别名') VARCHAR(40)"`
	Source       int    `json:"source" xorm:"not null default 1 comment('来源：1、APP注册；2、后台添加；') TINYINT(1)"`
	Status       int    `json:"status" xorm:"not null default 1 comment('是否启用 1、启用  2、停用') TINYINT(1)"`
	AppVersion   string `json:"appVersion" xorm:"default '''' comment('客户端版本号') VARCHAR(30)"`
	Code         string `json:"code" xorm:"default 'NULL' comment('我的推广码') VARCHAR(10)"`
	LoginIp      string `json:"loginIp" xorm:"default 'NULL' comment('最近登录IP') VARCHAR(30)"`
	LoginTime    int64  `json:"loginTime" xorm:"default 'NULL' comment('登录时间') DATETIME"`
	LoginRegion  string `json:"loginRegion" xorm:"default 'NULL' comment('上次登录地点') VARCHAR(20)"`
	LoginCount   int    `json:"loginCount" xorm:"not null default 0 comment('登录总次数') INT(10)"`
	CreateUser   int    `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime   int64  `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser   int    `json:"update_user" xorm:"not null default 0 comment('修改人') INT(11)"`
	UpdateTime   int64  `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int    `json:"mark" xorm:"not null default 1 comment('有效标识：1正常 0删除') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *Member) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *Member) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *Member) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *Member) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&Member{})
}

//批量删除
func (r *Member) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&Member{})
}
