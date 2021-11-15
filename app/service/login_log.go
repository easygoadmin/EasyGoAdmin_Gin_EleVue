/**
 *
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
