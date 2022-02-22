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
 * 操作日志-服务类
 * @author 半城风雨
 * @since 2021/11/12
 * @File : oper_log
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model"
	"easygoadmin/utils"
)

var OperLog = new(operLogService)

type operLogService struct{}

func (s *operLogService) GetList(req *dto.OperLogPageReq) ([]model.OperLog, int64, error) {
	// 初始化查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	if req != nil {
		// 操作账号
		if req.Username != "" {
			query = query.Where("username like ?", "%"+req.Username+"%")
		}
		// 操作模块
		if req.Model != "" {
			query = query.Where("model like ?", "%"+req.Model+"%")
		}
	}
	// 排序
	query = query.Desc("id")
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 实例化对象
	list := make([]model.OperLog, 0)
	count, err := query.FindAndCount(&list)
	return list, count, err
}
