package controllers

import "github.com/gin-gonic/gin"

type DemoController struct {
	BaseController
}

func (d *DemoController) Index(c *gin.Context) {
	d.RenderJson(c, 200, 0, "", map[string]string{"test": "yes"})
}
