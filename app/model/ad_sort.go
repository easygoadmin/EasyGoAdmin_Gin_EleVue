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

type AdSort struct {
	Id          int    `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Description string `json:"description" xorm:"not null comment('广告位描述') VARCHAR(255)"`
	ItemId      int    `json:"itemId" xorm:"not null default 0 comment('站点ID') INT(10)"`
	CateId      int    `json:"cateId" xorm:"not null default 0 comment('栏目ID') SMALLINT(5)"`
	LocId       int    `json:"locId" xorm:"not null default 0 comment('广告页面位置') SMALLINT(5)"`
	Platform    int    `json:"platform" xorm:"not null default 1 comment('站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端') TINYINT(1)"`
	Sort        int    `json:"sort" xorm:"not null default 125 comment('广告位排序') SMALLINT(5)"`
	CreateUser  int    `json:"create_user" xorm:"default 0 comment('添加人') INT(10)"`
	CreateTime  int64  `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser  int    `json:"update_user" xorm:"default 0 comment('更新人') INT(10)"`
	UpdateTime  int64  `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark        int    `json:"mark" xorm:"not null default 1 comment('有效标识(1正常 0删除)') TINYINT(1)"`
}

// 根据条件查询单条数据
func (r *AdSort) Get() (bool, error) {
	return utils.XormDb.Get(r)
}

// 插入数据
func (r *AdSort) Insert() (int64, error) {
	return utils.XormDb.Insert(r)
}

// 更新数据
func (r *AdSort) Update() (int64, error) {
	return utils.XormDb.Id(r.Id).Update(r)
}

// 删除
func (r *AdSort) Delete() (int64, error) {
	return utils.XormDb.Id(r.Id).Delete(&AdSort{})
}

//批量删除
func (r *AdSort) BatchDelete(ids ...int64) (int64, error) {
	return utils.XormDb.In("id", ids).Delete(&AdSort{})
}
