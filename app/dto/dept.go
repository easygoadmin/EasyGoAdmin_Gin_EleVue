/**
 *
 * @author 半城风雨
 * @since 2021/9/13
 * @File : dept
 */
package dto

import "easygoadmin/app/model"

// 分页查询条件
type DeptPageReq struct {
	Name  string `form:"name"`  // 部门名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加部门
type DeptAddReq struct {
	Name     string `form:"name" binding:"required"`
	Code     string `form:"code" binding:"required"`
	Fullname string `form:"fullname" binding:"required"`
	Type     int    `form:"type" binding:"required"`
	Pid      int    `form:"pid"`
	Sort     int    `form:"sort" binding:"required"`
	Note     string
}

// 部门编辑
type DeptUpdateReq struct {
	Id       int    `form:"id" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Code     string `form:"code" binding:"required"`
	Fullname string `form:"fullname" binding:"required"`
	Type     int    `form:"type" binding:"required"`
	Pid      int    `form:"pid"`
	Sort     int    `form:"sort" binding:"required"`
	Note     string
}

// 部门树结构
type DeptTreeNode struct {
	model.Dept
	Children []*DeptTreeNode `json:"children"` // 子栏目
}
