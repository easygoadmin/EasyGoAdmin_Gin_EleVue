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

type User struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	Realname     string `json:"realname" xorm:"default 'NULL' comment('真实姓名') index VARCHAR(150)"`
	Nickname     string `json:"nickname" xorm:"default 'NULL' comment('昵称') VARCHAR(150)"`
	Gender       int    `json:"gender" xorm:"default 3 comment('性别:1男 2女 3保密') TINYINT(1)"`
	Avatar       string `json:"avatar" xorm:"default 'NULL' comment('头像') VARCHAR(150)"`
	Mobile       string `json:"mobile" xorm:"default 'NULL' comment('手机号码') CHAR(11)"`
	Email        string `json:"email" xorm:"default 'NULL' comment('邮箱地址') VARCHAR(30)"`
	Birthday     int64  `json:"birthday" xorm:"default 'NULL' comment('出生日期') DATE"`
	DeptId       int    `json:"deptId" xorm:"default 0 comment('部门ID') INT(11)"`
	LevelId      int    `json:"levelId" xorm:"default 0 comment('职级ID') INT(11)"`
	PositionId   int    `json:"positionId" xorm:"default 0 comment('岗位ID') SMALLINT(3)"`
	ProvinceCode string `json:"province_code" xorm:"default 'NULL' comment('省份编号') VARCHAR(50)"`
	CityCode     string `json:"city_code" xorm:"default 'NULL' comment('市区编号') VARCHAR(50)"`
	DistrictCode string `json:"district_code" xorm:"default 'NULL' comment('区县编号') VARCHAR(50)"`
	Address      string `json:"address" xorm:"default 'NULL' comment('详细地址') VARCHAR(255)"`
	CityName     string `json:"city_name" xorm:"default 'NULL' comment('所属城市') VARCHAR(150)"`
	Username     string `json:"username" xorm:"default 'NULL' comment('登录用户名') VARCHAR(50)"`
	Password     string `json:"password" xorm:"default 'NULL' comment('登录密码') VARCHAR(150)"`
	Salt         string `json:"salt" xorm:"default 'NULL' comment('盐加密') VARCHAR(30)"`
	Intro        string `json:"intro" xorm:"default 'NULL' comment('个人简介') VARCHAR(500)"`
	Status       int    `json:"status" xorm:"default 1 comment('状态：1正常 2禁用') TINYINT(1)"`
	Note         string `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(500)"`
	Sort         int    `json:"sort" xorm:"default 125 comment('排序号') INT(11)"`
	LoginNum     int    `json:"login_num" xorm:"default 0 comment('登录次数') INT(11)"`
	LoginIp      string `json:"login_ip" xorm:"default 'NULL' comment('最近登录IP') VARCHAR(20)"`
	LoginTime    int64  `json:"login_time" xorm:"default 'NULL' comment('最近登录时间') DATETIME"`
	CreateUser   int    `json:"create_user" xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime   int64  `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME"`
	UpdateUser   int    `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime   int64  `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark         int    `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *User) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *User) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *User) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *User) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&User{})
}

//批量删除
func (r *User) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&User{})
}
