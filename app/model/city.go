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

type City struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('编号') INT(11)"`
	Pid        int    `json:"pid" xorm:"not null default 0 comment('父级编号') INT(11)"`
	Level      int    `json:"level" xorm:"not null default 0 comment('城市级别：1省 2市 3区') TINYINT(1)"`
	Name       string `json:"name" xorm:"not null comment('城市名称') index VARCHAR(50)"`
	Citycode   string `json:"citycode" xorm:"not null comment('城市编号（区号）') VARCHAR(10)"`
	PAdcode    string `json:"PAdcode" xorm:"default 'NULL' comment('父级地理编号') VARCHAR(10)"`
	Adcode     string `json:"adcode" xorm:"default 'NULL' comment('地理编号') VARCHAR(10)"`
	Lng        string `json:"lng" xorm:"default 'NULL' comment('城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7') VARCHAR(10)""`
	Lat        string `json:"lat" xorm:"default 'NULL' comment('城市坐标中心点纬度（* 1e6）') VARCHAR(10)""`
	Sort       int    `json:"sort" xorm:"not null default 125 comment('排序号') TINYINT(3)"`
	CreateUser int    `json:"create_user" xorm:"not null default 0 comment('添加人') INT(11)"`
	CreateTime int64  `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME"`
	UpdateUser int    `json:"update_user" xorm:"not null default 0 comment('更新人') INT(11)"`
	UpdateTime int64  `json:"update_time" xorm:"default 'NULL' comment('更新时间') DATETIME"`
	Mark       int    `json:"mark" xorm:"not null default 1 comment('有效标记') TINYINT(1)"`
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
