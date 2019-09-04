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

	// 文章操作
	{
		// 获取用户所有文章
		authorized.GET("", handler.GetUserFiles)
		// 创建新文章
		authorized.POST("", handler.CreateFile)
		// 删除文章
		authorized.DELETE("", handler.DeleteFile)
	}

	r.Run(":49158")
}
