/**
 *
 * @author 摆渡人
 * @since 2021/9/9
 * @File : menu
 */
package service

import (
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
)

// 中间件管理服务
var Menu = new(menuService)

type menuService struct{}

// 获取菜单权限列表
func (s *menuService) GetPermissionList(userId int) interface{} {
	if userId == 1 {
		// 管理员(拥有全部权限)
		menuList, _ := Menu.GetTreeList()
		return menuList
	} else {
		//// 非管理员
		//// 创建查询实例
		//query := dao.Menu.As("m").Clone()
		//// 内联查询
		//query = query.InnerJoin("sys_role_menu as r", "m.id = r.menu_id")
		//query = query.InnerJoin("sys_user_role ur", "ur.role_id=r.role_id")
		//query = query.Where("ur.user_id=? AND m.type=0 AND m.`status`=1 AND m.mark=1", userId)
		//// 获取字段
		//query.Fields("m.*")
		//// 排序
		//query = query.Order("m.id asc")
		//// 数据转换
		//var list []*model.Menu
		//query.Structs(&list)
		//// 数据处理
		//var menuNode model.TreeNode
		//makeTree(list, &menuNode)
		//return menuNode.Children
		return nil
	}
}

// 获取子级菜单
func (s *menuService) GetTreeList() ([]*vo.MenuTreeNode, error) {
	var menuNode vo.MenuTreeNode
	list := make([]model.Menu, 0)
	err := utils.XormDb.Where("type=0 and mark=1").OrderBy("sort").Find(&list)
	if err != nil {
		return nil, err
	}
	makeTree(list, &menuNode)
	return menuNode.Children, nil
}

//递归生成分类列表
func makeTree(menu []model.Menu, tn *vo.MenuTreeNode) {
	for _, c := range menu {
		if c.ParentId == tn.Id {
			child := &vo.MenuTreeNode{}
			child.Menu = c
			tn.Children = append(tn.Children, child)
			makeTree(menu, child)
		}
	}
}
