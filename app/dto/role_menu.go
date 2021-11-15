/**
 *
 * @author 半城风雨
 * @since 2021/9/14
 * @File : role_menu
 */
package dto

// 角色菜单数据
type RoleMenuSaveReq struct {
	RoleId  int   `form:"roleId" binding:"required"`
	MenuIds []int `form:"menuIds" binding:"required"`
}
