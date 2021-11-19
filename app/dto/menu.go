// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 菜单Dto
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
