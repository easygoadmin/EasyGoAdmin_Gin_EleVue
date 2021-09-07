/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : level
 */
package service

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/model/level"
	"easygoadmin/library/db"
	"easygoadmin/utils/convert"
	"easygoadmin/utils/gconv"
	"errors"
	"github.com/go-xorm/xorm"
	"time"
)

var x *xorm.Engine

// 中间件管理服务
var Level = new(levelService)

type levelService struct{}

func (s *levelService) GetList(req *dto.LevelPageReq) ([]level.Level, int64, error) {
	// 实例化引擎
	db := db.Instance().Engine()
	if db == nil {
		return nil, 0, errors.New("连接数据库失败")
	}
	// 初始化查询实例
	query := db.Table(level.TableName()).Where("mark=1")
	if req != nil {
		// 职级名称查询
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}

	// 查询记录总数
	totalQuery := query.Clone()
	count, err := totalQuery.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Asc("sort")
	// 分页
	query = query.Limit(req.Limit, 0)
	// 对象转换
	var list []level.Level
	query.Find(&list)
	return list, count, nil
}

func (s *levelService) Add(req *dto.LevelAddReq, userId int) (int, error) {
	// 实例化对象
	var entity level.Level
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	_, err := entity.Insert()
	if err != nil {
		return 0, err
	}

	return entity.Id, nil
}

func (s *levelService) Update(req *dto.LevelUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &level.Level{Id: req.Id}
	isOk, err := entity.FindOne()
	if !isOk || err != nil {
		return 0, err
	}
	if entity == nil {
		return 0, errors.New("记录不存在")
	}
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

// 删除
func (s *levelService) Delete(Ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(Ids, ",")
	if len(idsArr) > 1 {
		// 单个删除
		entity := &level.Level{Id: gconv.Int(Ids)}
		rows, err := entity.Delete()
		if err != nil {
			return 0, err
		}
		if rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		return 1, nil
	}
}

func (s *levelService) Status(req *dto.LevelStatusReq, userId int) (int64, error) {
	entity := &level.Level{Id: req.Id}
	isOk, err := entity.FindOne()
	if !isOk || err != nil {
		return 0, err
	}
	if entity == nil {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}
