package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"resmgt/handler"
	"resmgt/util"
)

// todo: 整理一些配置到 main 中
// todo: 分类管理
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
		authorized.GET("", handler.AuthGet, handler.GetUserFiles)
		// 创建新文章
		authorized.POST("", handler.AuthCreate, handler.CreateFile)
		// 删除文章
		authorized.DELETE("", handler.AuthDelete, handler.DeleteFile)
	}

	if err := endless.ListenAndServe(":49158", r); err != nil {
		// 释放资源
		util.MyDB.Close()
	}
}
