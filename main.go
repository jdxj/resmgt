package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"resmgt/handler"
	"resmgt/util"
)

// todo: 整理一些配置到 main 中
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

	// 文章操作
	filesAuth := r.Group("/files")
	filesAuth.Use(handler.Authenticate)
	{
		// 获取用户所有文章
		filesAuth.GET("", handler.AuthGet, handler.GetUserFiles)
		// 创建新文章
		filesAuth.POST("", handler.AuthCreate, handler.CreateFile)
		// 删除文章
		filesAuth.DELETE("", handler.AuthDelete, handler.DeleteFile)
	}

	catAuth := r.Group("/categories")
	catAuth.Use(handler.Authenticate)
	// 分类操作
	{
		// 获取分类
		catAuth.GET("", handler.AuthCateBrowse, handler.GetUserCategories)
		// 创建分类
		catAuth.POST("", handler.AuthCateCreate, handler.CreateCategory)
		// 删除分类
		catAuth.DELETE("", handler.AuthCateDelete, handler.DeleteCategory)
	}

	if err := endless.ListenAndServe(":49158", r); err != nil {
		// 释放资源
		util.MyDB.Close()
	}
}
