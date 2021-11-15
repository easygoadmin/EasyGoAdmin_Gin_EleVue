/**
 *
 * @author 半城风雨
 * @since 2021/11/12
 * @File : user_role
 */
package service

import (
	"easygoadmin/app/model"
	"easygoadmin/utils"
)

var UserRole = new(userRoleService)

type userRoleService struct{}

// 获取用户角色列表
func (s *userRoleService) GetUserRoleList(userId int) []model.Role {
	// 实例化对象
	list := make([]model.Role, 0)
	utils.XormDb.Table("sys_role").Alias("r").
		Join("INNER", []string{"sys_user_role", "ur"}, "r.id=ur.role_id").
		Where("ur.user_id=? AND r.mark=1", userId).
		Cols("r.*").
		OrderBy("r.sort asc").
		Find(&list)
	return list
}
