/**
 *
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
	Sort     int    `form:"sort"        binding:"required"` // 排序
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
	Sort     int    `form:"sort"        binding:"required"` // 排序
	Note     string `form:"note"`                           // 配置说明
}

// 删除字典项
type ConfigDataDeleteReq struct {
	Ids string `form:"ids" v:"required#请选择需要删除的数据记录"`
}

// 设置状态
type ConfigDataStatusReq struct {
	Id     int `form:"id" v:"required#主键ID不能为空"`
	Status int `form:"status"    v:"required#状态不能为空"`
}
