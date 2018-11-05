package routes

import (
	"GinWeb/app/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type api struct {
}

func (api api) Run() {

	demo := new(controllers.DemoController)

	R.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// 设置游戏数据库连接
	apiRouter := R.Group("/").Use()
	{
		apiRouter.GET("/demo/index", demo.Index)
	}

	//R.StaticFile("schema.json", "./api_schema.json")
	R.GET("health", health)
	R.POST("health", health)

}

func health(c *gin.Context) {
	c.String(200, "ok")
}
