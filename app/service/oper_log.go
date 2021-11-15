/**
 *
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
