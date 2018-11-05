package middleware

import (
	"GinWeb/app/helpers"
	"github.com/gin-gonic/gin"
)

// 赋值请求参数到全局变量
func RequestParam(c *gin.Context) {
	c.Request.ParseForm()
	for k, v := range c.Request.Form {
		// 先默认将第一个赋值到全局变量，以后有单参数有多值的，再进行判断后赋值
		helpers.RequestParam[k] = v[0]
	}
	c.Next()
}
