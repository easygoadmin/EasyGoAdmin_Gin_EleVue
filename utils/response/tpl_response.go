package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

// 返回一个tpl响应
func BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

// 返回一个错误的tpl响应
func ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/error.html",
	}
	return &t
}

// 输出页面模板附加自定义函数
func (resp *TplResp) WriteTpl(params ...gin.H) {
	if params == nil || len(params) == 0 {
		fmt.Println("渲染")
		fmt.Println(resp.tpl)
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{"uid": 1})
	} else {
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
}
