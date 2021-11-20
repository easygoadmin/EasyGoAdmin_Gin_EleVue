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
 * 部门-服务类
 * @author 半城风雨
 * @since 2021/9/13
 * @File : dept
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strings"
	"time"
)

var Dept = new(deptService)

type deptService struct{}

func (s *deptService) GetList(req *dto.DeptPageReq) ([]model.Dept, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 部门名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []model.Dept
	err := query.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *deptService) Add(req *dto.DeptAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Dept
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.Mark = 1
	// 插入记录
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, errors.New("添加失败")
	}
	return rows, nil
}

func (s *deptService) Update(req *dto.DeptUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &model.Dept{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}
	// 设置参数
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新记录
	rows, err := entity.Update()
	if err != nil || rows == 0 {
		return 0, err
	}
	return rows, nil
}

func (s *deptService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Dept{Id: gconv.Int(ids)}
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
