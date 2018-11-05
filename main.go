package main

import (
	mw "GinWeb/app/middleware"
	"GinWeb/config"
	"GinWeb/routes"
	"github.com/gin-gonic/gin"
)

func loadRoute() {
	if config.C.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化引擎
	routes.R = gin.Default()
	middleware()
	routes.Init()
	routes.R.Run(config.C.Http.Listen)
}

func middleware() {
	//routes.R.Use(mw.ValidateTokenMiddleware)
	routes.R.Use(mw.RequestParam)
	//routes.R.Use(mw.ValidateCommonData)
	routes.R.Use(mw.CORSMiddleware())
}

func main() {
	loadRoute()
}
