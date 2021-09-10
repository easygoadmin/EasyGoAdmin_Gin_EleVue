/**
 * mysql数据库配置信息
 */
package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool // 是否正常运行
}

// 系统中所有mysql主库 root:root@tcp(127.0.0.1:3306)/lottery?charset=utf-8
var DbMasterList = []DbConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "",
		Database:  "easygoadmin.gin.ele",
		IsRunning: true,
	},
}

var DbMaster DbConfig = DbMasterList[0]
