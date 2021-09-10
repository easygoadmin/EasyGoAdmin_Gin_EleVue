/**
 *
 * @author 摆渡人
 * @since 2021/9/8
 * @File : db
 */
package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

//引擎对象
var DbEngine *xorm.Engine
var err error

// 初始化连接数据库
func init() {
	fmt.Println("初始化连接数据库")

	// 获取配置实例
	//config := cfg.Instance()
	// 创建引擎(此时还未连接数据库)
	DbEngine, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/easygoadmin.gin.ele?charset=utf8")
	if err != nil {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}
	// 通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库。
	err = DbEngine.Ping()
	if err == nil {
		fmt.Println("数据库连接成功")
		////关闭连接
		//defer engine.Close()
	} else {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	DbEngine.SetTableMapper(tbMapper)
	//// 开启调试模式和打印日志
	//if cfg.Instance().Database.Debug {
	//	engine.ShowSQL(cfg.Instance().Database.Debug)
	//	engine.Logger().SetLevel(core.LOG_DEBUG)
	//}
	DbEngine.ShowSQL(true)
	DbEngine.Logger().SetLevel(core.LOG_DEBUG)

	//level := &model.Level{}
	//level.Name = "测试22"
	//level.Status = 1
	//level.Sort = 1
	//level.CreateUser = 1
	//level.CreateTime = time.Now()
	//level.Mark = 1
	//rows, err := engine.Insert(level)
	//fmt.Println(rows)
	//fmt.Println(err)
	//fmt.Println("11")
}
