package main

import (
	"github.com/gin-gonic/gin"
	"resmgt/handler"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.Home)

	// login
	{
		r.GET("/login", func(c *gin.Context) {
			c.JSON(400, gin.H{
				"err": "please use post method",
			})
		})
		r.POST("/login", handler.Login)
	}

	authorized := r.Group("/files")
	authorized.Use(handler.Authenticate)
	{
		authorized.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"state": "pass",
			})
		})
	}

	r.Run(":49158")
}
