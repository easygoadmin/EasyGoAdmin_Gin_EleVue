/**
 *
 * @author 半城风雨
 * @since 2021/9/9
 * @File : menu
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gstr"
	"errors"
	"strings"
	"time"
)

// 中间件管理服务
var Menu = new(menuService)

type menuService struct{}

func (s *menuService) GetList(req *dto.MenuQueryReq) ([]model.Menu, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 部门名称
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询列表
	var list []model.Menu
	err := query.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *menuService) Add(req *dto.MenuAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Menu
	entity.ParentId = req.ParentId
	entity.Title = req.Title
	entity.Icon = req.Icon
	entity.Path = req.Path
	entity.Component = req.Component
	entity.Target = req.Target
	entity.Permission = req.Permission
	entity.Type = req.Type
	entity.Status = req.Status
	entity.Hide = req.Hide
	entity.Note = req.Note
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	// 添加节点
	setPermission(req.Type, req.CheckedList, req.Title, req.Path, gconv.Int(entity.Id), userId)
	return rows, nil
}

func (s *menuService) Update(req *dto.MenuUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Menu{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	entity.ParentId = req.ParentId
	entity.Title = req.Title
	entity.Icon = req.Icon
	entity.Path = req.Path
	entity.Component = req.Component
	entity.Target = req.Target
	entity.Permission = req.Permission
	entity.Type = req.Type
	entity.Status = req.Status
	entity.Hide = req.Hide
	entity.Note = req.Note
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新数据
	rows, err := entity.Update()
	if err != nil || rows == 0 {
		return 0, errors.New("更新失败")
	}

	// 添加节点
	setPermission(req.Type, req.CheckedList, req.Title, req.Path, req.Id, userId)

	return rows, nil
}

func (s *menuService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Menu{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, err
		}
		return rows, nil
	} else {
		// 批量删除
		return 0, nil
	}
}

// 添加节点
func setPermission(menuType int, checkedList []int, name string, url string, parentId int, userId int) {
	if menuType != 0 || len(checkedList) == 0 || url == "" {
		return
	}
	// 删除现有节点
	utils.XormDb.Where("parent_id=?", parentId).Delete(&model.Menu{})
	// 模块名称
	moduleTitle := gstr.Replace(name, "管理", "")
	// 创建权限节点
	urlArr := strings.Split(url, "/")

	if len(urlArr) >= 3 {
		// 模块名
		moduleName := urlArr[len(urlArr)-1]
		// 节点处理
		for _, v := range checkedList {
			// 实例化对象
			var entity model.Menu
			// 节点索引
			value := gconv.Int(v)
			if value == 1 {
				entity.Title = "查询" + moduleTitle
				entity.Path = "/" + moduleName + "/list"
				entity.Permission = "sys:" + moduleName + ":list"
				entity.Method = "GET"
			} else if value == 5 {
				entity.Title = "添加" + moduleTitle
				entity.Path = "/" + moduleName + "/add"
				entity.Permission = "sys:" + moduleName + ":add"
				entity.Method = "POST"
			} else if value == 10 {
				entity.Title = "修改" + moduleTitle
				entity.Path = "/" + moduleName + "/update"
				entity.Permission = "sys:" + moduleName + ":update"
				entity.Method = "PUT"
			} else if value == 15 {
				entity.Title = "删除" + moduleTitle
				entity.Path = "/" + moduleName + "/delete"
				entity.Permission = "sys:" + moduleName + ":delete"
				entity.Method = "DELETE"
			} else if value == 20 {
				entity.Title = moduleTitle + "详情"
				entity.Path = "/" + moduleName + "/detail"
				entity.Permission = "sys:" + moduleName + ":detail"
				entity.Method = "GET"
			} else if value == 25 {
				entity.Title = "设置状态"
				entity.Path = "/" + moduleName + "/status"
				entity.Permission = "sys:" + moduleName + ":status"
				entity.Method = "PUT"
			} else if value == 30 {
				entity.Title = "批量删除"
				entity.Path = "/" + moduleName + "/dall"
				entity.Permission = "sys:" + moduleName + ":dall"
				entity.Method = "DELETE"
			} else if value == 35 {
				entity.Title = "添加子级"
				entity.Path = "/" + moduleName + "/addz"
				entity.Permission = "sys:" + moduleName + ":addz"
				entity.Method = "POST"
			} else if value == 40 {
				entity.Title = "全部展开"
				entity.Path = "/" + moduleName + "/expand"
				entity.Permission = "sys:" + moduleName + ":expand"
				entity.Method = "GET"
			} else if value == 45 {
				entity.Title = "全部折叠"
				entity.Path = "/" + moduleName + "/collapse"
				entity.Permission = "sys:" + moduleName + ":collapse"
				entity.Method = "GET"
			} else if value == 50 {
				entity.Title = "导出" + moduleTitle
				entity.Path = "/" + moduleName + "/export"
				entity.Permission = "sys:" + moduleName + ":export"
				entity.Method = "GET"
			} else if value == 55 {
				entity.Title = "导入" + moduleTitle
				entity.Path = "/" + moduleName + "/import"
				entity.Permission = "sys:" + moduleName + ":import"
				entity.Method = "GET"
			} else if value == 60 {
				entity.Title = "分配权限"
				entity.Path = "/" + moduleName + "/permission"
				entity.Permission = "sys:" + moduleName + ":permission"
				entity.Method = "POST"
			} else if value == 65 {
				entity.Title = "重置密码"
				entity.Path = "/" + moduleName + "/resetPwd"
				entity.Permission = "sys:" + moduleName + ":resetPwd"
				entity.Method = "PUT"
			}
			entity.ParentId = parentId
			entity.Type = 1
			entity.Status = 1
			entity.Target = "_self"
			entity.Sort = value
			entity.CreateUser = userId
			entity.CreateTime = time.Now()
			entity.UpdateUser = userId
			entity.UpdateTime = time.Now()
			entity.Mark = 1

			// 插入节点
			entity.Insert()
		}
	}
}

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
