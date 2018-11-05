package controllers

import (
	"github.com/gin-gonic/gin"
)

func init() {
	b := &BaseController{}
	b.init()
}

type BaseController struct {
}

func (bc *BaseController) init() {
}

func (bc *BaseController) RenderJson(c *gin.Context, httpCode int, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	c.JSON(httpCode, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
	c.Abort()
}
