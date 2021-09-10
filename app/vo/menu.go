/**
 *
 * @author 摆渡人
 * @since 2021/9/9
 * @File : menu
 */
package vo

import "easygoadmin/app/model"

// 菜单Vo
type MenuTreeNode struct {
	model.Menu
	Children []*MenuTreeNode `json:"children"` // 子菜单
}
