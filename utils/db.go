/**
 *
 * @author 摆渡人
 * @since 2021/9/8
 * @File : db
 */
package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var XormDb *xorm.Engine

func init() {
	fmt.Println("初始化并连接数据库")
	// 1.连接数据库
	sqlStr := "root:@tcp(127.0.0.1:3306)/easygoadmin.gin.ele?charset=utf8&parseTime=true&loc=Local"
	var err error
	XormDb, err = xorm.NewEngine("mysql", sqlStr)
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}

	// 结构体与数据表的映射
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	XormDb.SetTableMapper(tbMapper)

	// 会在控制台打印执行的sql
	XormDb.ShowSQL(true)
	XormDb.Logger().SetLevel(core.LOG_DEBUG)
}
