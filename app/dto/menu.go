/**
 *
 * @author 半城风雨
 * @since 2021/9/13
 * @File : menu
 */
package dto

import "easygoadmin/app/model"

// 列表查询条件
type MenuQueryReq struct {
	Title string `form:"name"` // 菜单标题
}

// 添加菜单
type MenuAddReq struct {
	ParentId    int    `form:"parent_id"`                      // 上级ID
	Title       string `form:"title"       binding:"required"` // 菜单标题
	Icon        string `form:"icon"        binding:"required"` // 图标
	Path        string `form:"path"        binding:"required"` // URL地址
	Component   string `form:"component"`                      // 菜单组件
	Target      string `form:"target"`                         // 打开方式：0组件 1内链 2外链
	Permission  string `form:"permission"`                     // 权限标识
	Type        int    `form:"type"`                           // 类型：1模块 2导航 3菜单 4节点
	Status      int    `form:"status"      binding:"required"` // 状态：1正常 2禁用
	Hide        int    `form:"hide"`                           // 是否可见：1是 2否
	Note        string `form:"note"`                           // 菜单备注
	Sort        int    `form:"sort"        binding:"required"` // 显示顺序
	CheckedList []int  `form:"checkedList"`                    // 权限节点
}

// 更新菜单
type MenuUpdateReq struct {
	Id          int    `form:"id" 		   binding:"required"`
	ParentId    int    `form:"parent_id"`                      // 上级ID
	Title       string `form:"title"       binding:"required"` // 菜单标题
	Icon        string `form:"icon"        binding:"required"` // 图标
	Path        string `form:"path"        binding:"required"` // URL地址
	Component   string `form:"component"`                      // 菜单组件
	Target      string `form:"target"`                         // 打开方式：0组件 1内链 2外链
	Permission  string `form:"permission"`                     // 权限标识
	Type        int    `form:"type"`                           // 类型：1模块 2导航 3菜单 4节点
	Status      int    `form:"status"      binding:"required"` // 是否显示：1显示 2不显示
	Hide        int    `form:"hide"`                           // 是否可见：1是 2否
	Note        string `form:"note"`                           // 菜单备注
	Sort        int    `form:"sort"        binding:"required"` // 显示顺序
	CheckedList []int  `form:"checkedList"`                    // 权限节点
}

// 菜单信息
type MenuInfoVo struct {
	model.Menu
	CheckedList []int `json:"checkedList"` // 权限节点列表
}
