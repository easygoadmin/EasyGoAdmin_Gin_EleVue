/**
 *
 * @author 半城风雨
 * @since 2021/11/13
 * @File : config_data
 */
package vo

import "easygoadmin/app/model"

// 配置数据列表
type ConfigDataVo struct {
	model.ConfigData
	TypeName string `json:"typeName"`
}
