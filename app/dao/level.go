/**
 *
 * @author 摆渡人
 * @since 2021/9/2
 * @File : levelDao
 */
package dao

import (
	"easygoadmin/app/model"
	"easygoadmin/library"
	"easygoadmin/utils"
)

var Level = new(levelDao)

type levelDao struct{}

// 根据结构体中已有的非空数据来获得单条数据
func (d *levelDao) GetById(Id int) (*model.Level, error) {
	level := &model.Level{Id: Id}
	isOk, err := library.Instance().Engine().Get(&level)
	if !isOk || err != nil {
		return nil, err
	}
	return level, nil
}

// 插入数据
func (d *levelDao) Insert(entity *model.Level) (int64, error) {
	return utils.XormDb.Insert(&entity)
}

// 更新数据
func (d *levelDao) Update(entity *model.Level) (int64, error) {
	return utils.XormDb.Id(entity.Id).Update(&entity)
}

// 删除
func (d *levelDao) Delete(entity *model.Level) (int64, error) {
	return utils.XormDb.Id(entity.Id).Delete(&model.Level{})
}
