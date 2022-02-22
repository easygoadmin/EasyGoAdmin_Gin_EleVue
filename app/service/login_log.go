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
 * 登录日志-服务类
 * @author 半城风雨
 * @since 2021/11/12
 * @File : login_log
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strings"
)

var LoginLog = new(loginLogService)

type loginLogService struct{}

func (s *loginLogService) GetList(req *dto.LoginLogPageReq) ([]model.LoginLog, int64, error) {
	// 初始化查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 操作用户
		if req.Username != "" {
			query = query.Where("username like ?", "%"+req.Username+"%")
		}
	}
	// 排序
	query = query.OrderBy("id desc")
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 实例化对象
	list := make([]model.LoginLog, 0)
	count, err := query.FindAndCount(&list)
	return list, count, err
}

func (s *loginLogService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.LoginLog{Id: gconv.Int64(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		return 0, nil
	}
}
