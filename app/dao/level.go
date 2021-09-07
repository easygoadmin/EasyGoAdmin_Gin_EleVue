/**
 *
 * @author 摆渡人
 * @since 2021/9/2
 * @File : levelDao
 */
package dao

import (
	"easygoadmin/app/model/level"
	"easygoadmin/library/db"
)

var Level = new(levelDao)

type levelDao struct{}

// 根据结构体中已有的非空数据来获得单条数据
func (d *levelDao) GetById(Id int) (*level.Level, error) {
	level := &level.Level{Id: Id}
	isOk, err := db.Instance().Engine().Get(&level)
	if !isOk || err != nil {
		return nil, err
	}
	return level, nil
}
//
//// 插入数据
//func (d *levelDao) Insert(entity *model.Level) (int64, error) {
//	return db.Instance().Engine().Insert(&entity)
//}
//
//// 更新数据
//func (d *levelDao) Update(entity *model.Level) (int64, error) {
//	return db.Instance().Engine().ID(entity.Id).Update(&entity)
//}
//
//// 删除
//func (d *levelDao) Delete(entity *model.Level) (int64, error) {
//	return db.Instance().Engine().ID(entity.Id).Delete(&entity)
//}
