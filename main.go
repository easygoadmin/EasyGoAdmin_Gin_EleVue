/**
 *
 * @author 摆渡人
 * @since 2021/8/20
 * @File : main
 */
package main

import (
	_ "easygoadmin/boot"
	"easygoadmin/library/cfg"
	_ "easygoadmin/router"
	_ "easygoadmin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 开始调试模式
	gin.SetMode("debug")

	// 实例化配置
	config := cfg.Instance()
	if config == nil {
		fmt.Printf("参数错误")
		return
	}

	//dataSourceName格式：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	//DbEngine.SetMapper(core.SameMapper{})
	//
	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	//DbEngine.SetTableMapper(tbMapper)

}
