/**
 *
 * @author 半城风雨
 * @since 2021/9/14
 * @File : role_menu
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
)

var RoleMenu = new(roleMenuService)

type roleMenuService struct{}

func (s *roleMenuService) GetRoleMenuList(roleId int) ([]vo.RoleMenuInfo, error) {
	// 获取全部菜单列表
	var menuList []model.Menu
	utils.XormDb.Where("status=1").Where("mark=1").OrderBy("sort asc").Find(&menuList)
	if len(menuList) == 0 {
		return nil, errors.New("菜单列表不存在")
	}
	// 获取角色菜单权限列表
	var roleMenuList []model.RoleMenu
	utils.XormDb.Where("role_id=?", roleId).Find(&roleMenuList)
	idList := make([]interface{}, 0)
	for _, v := range roleMenuList {
		idList = append(idList, v.MenuId)
	}

	// 对象处理
	var list []vo.RoleMenuInfo
	if len(menuList) > 0 {
		for _, m := range menuList {
			var info vo.RoleMenuInfo
			info.Id = m.Id
			info.Title = m.Title
			info.Open = true
			info.ParentId = m.ParentId
			// 节点选中值
			if utils.InArray(gconv.String(m.Id), idList) {
				info.Checked = true
			}
			list = append(list, info)
		}
	}
	return list, nil
}

func (s *roleMenuService) Save(req *dto.RoleMenuSaveReq) error {
	if utils.AppDebug() {
		return errors.New("演示环境，暂无权限操作")
	}
	itemArr := req.MenuIds
	if len(itemArr) == 0 {
		return errors.New("请选择权限节点")
	}
	// 删除现有的角色权限数据
	utils.XormDb.Where("role_id=?", req.RoleId).Delete(&model.RoleMenu{})
	// 遍历创建新角色权限数据
	for _, v := range itemArr {
		var entity model.RoleMenu
		entity.RoleId = req.RoleId
		entity.MenuId = v
		entity.Insert()
	}
	return nil
}
