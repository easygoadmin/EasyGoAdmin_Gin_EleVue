/**
 *
 * @author 半城风雨
 * @since 2021/9/9
 * @File : login
 */
package dto

// 系统登录
type LoginReq struct {
	UserName string `form:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" binding:"required,min=6,max=12"`
	Captcha  string `form:"captcha" binding:"required,min=4,max=6"`
	IdKey    string `form:"idKey" binding:"required"`
}
