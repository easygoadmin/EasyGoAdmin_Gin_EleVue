/**
 *
 * @author 半城风雨
 * @since 2021/9/14
 * @File : role_menu
 */
package vo

// 角色权限菜单列表
type RoleMenuInfo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ParentId int    `json:"parentId"`
	Checked  bool   `json:"checked"`
	Open     bool   `json:"open"`
}
