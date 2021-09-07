// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

/**
 * 转换函数库
 * @author 半城风雨
 * @since 2021/7/14
 * @File : convert
 */
package convert

import (
	"easygoadmin/utils/gconv"
	"strings"
)

// 带分隔符的字符串分裂成int64数组
func ToInt64Array(str, split string) []int64 {
	result := make([]int64, 0)
	if str == "" {
		return result
	}
	arr := strings.Split(str, split)
	if len(arr) > 0 {
		for i := range arr {
			if arr[i] != "" {
				result = append(result, gconv.Int64(arr[i]))
			}
		}
	}
	return result

}
