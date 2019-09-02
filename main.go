package main

import (
	"github.com/gin-gonic/gin"
	"resmgt/handler"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.Home)

	// todo: 权限验证 (登录)
	authorized := r.Group("/auth", gin.BasicAuth(gin.Accounts{
		"jdxj": "jdxj",
		"test": "test",
	}))
	authorized.GET("/secrets", handler.Login)

	r.Run(":49158")
}
