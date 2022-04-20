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

/**
 * 配置数据Dto
 * @author 半城风雨
 * @since 2021/11/13
 * @File : config_data
 */
package dto

// 字典项列表查询条件
type ConfigDataPageReq struct {
	ConfigId int    `form:"configId"` // 字典ID
	Title    string `form:"name"`     // 配置标题
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加字典项
type ConfigDataAddReq struct {
	Title    string `form:"title"       binding:"required"` // 配置标题
	Code     string `form:"code"        binding:"required"` // 配置编码
	Value    string `form:"value"`                          // 配置值
	Options  string `form:"options"`                        // 配置项
	ConfigId int    `form:"config_id"   binding:"required"` // 配置ID
	Type     string `form:"type"        binding:"required"` // 配置类型
	Sort     int    `form:"sort"`                           // 排序
	Note     string `form:"note"`                           // 配置说明
}

// 更新字典项
type ConfigDataUpdateReq struct {
	Id       int    `form:"id" binding:"required"`
	Title    string `form:"title"       binding:"required"` // 配置标题
	Code     string `form:"code"        binding:"required"` // 配置编码
	Value    string `form:"value"`                          // 配置值
	Options  string `form:"options"`                        // 配置项
	ConfigId int    `form:"config_id"   binding:"required"` // 配置ID
	Type     string `form:"type"        binding:"required"` // 配置类型
	Sort     int    `form:"sort"`                           // 排序
	Note     string `form:"note"`                           // 配置说明
}

// 设置状态
type ConfigDataStatusReq struct {
	Id     int `form:"id" binding:"required"`
	Status int `form:"status"    binding:"required"`
}
