package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ValidateTokenMiddleware(c *gin.Context) {
	c.Request.ParseForm()
	for k, v := range c.Request.PostForm {
		fmt.Printf("k:%v\n", k)
		fmt.Printf("v:%v\n", v)
	}
	return
}
