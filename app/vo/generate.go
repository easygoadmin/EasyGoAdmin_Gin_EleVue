/**
 *
 * @author 半城风雨
 * @since 2021/11/15
 * @File : generate
 */
package vo

import "time"

// 数据库信息
type GenerateInfo struct {
	Name          string    `json:"name"`           // 表名
	Engine        string    `json:"engine"`         // 引擎
	Version       string    `json:"version"`        // 版本
	Collation     string    `json:"collation"`      // 编码
	Rows          int       `json:"rows"`           // 记录数
	DataLength    int       `json:"data_length"`    // 大小
	AutoIncrement int       `json:"auto_increment"` // 自增索引
	Comment       string    `json:"comment"`        // 表备注
	CreateTime    time.Time `json:"createTime"`     // 添加时间
	UpdateTime    time.Time `json:"updateTime"`     // 更新时间
}
