package routes

import "github.com/gin-gonic/gin"

type Routes interface {
	Run()
}

var R *gin.Engine

func Init() {
	var r Routes

	r = new(api)
	r.Run()
}
