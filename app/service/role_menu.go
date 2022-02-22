// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
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
 * 角色菜单-服务类
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
