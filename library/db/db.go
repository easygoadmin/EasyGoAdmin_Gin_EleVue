/**
 *
 * @author 半城风雨
 * @since 2021/9/7
 * @File : db
 */
package db

import (
	"easygoadmin/library/cfg"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

//引擎对象
var engine *xorm.Engine
var err error

// 初始化连接数据库
func init() {
	fmt.Println("初始化连接数据库")

	// 获取配置实例
	config := cfg.Instance()
	// 创建引擎(此时还未连接数据库)
	engine, err := xorm.NewEngine("mysql", config.Database.Master)
	if err != nil {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}
	// 通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库。
	err = engine.Ping()
	if err == nil {
		fmt.Println("数据库连接成功")
		//关闭连接
		defer engine.Close()
	} else {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}
	// 开启调试模式和打印日志
	if cfg.Instance().Database.Debug {
		engine.ShowSQL(cfg.Instance().Database.Debug)
		engine.Logger().SetLevel(core.LOG_DEBUG)
	}
}
