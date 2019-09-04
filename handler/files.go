package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"resmgt/module"
	"resmgt/util"
	"strconv"
	"time"
)

func GetUserFiles(c *gin.Context) {
	// todo: 获取文章权限检查
	token, err := c.Cookie("id")
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": "not found cookie id",
		})
		return
	}
	item, err := util.GetCache().Get(token)
	if item == nil || err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": "not found user",
		})
		return
	}

	var user module.User
	var files []module.File
	json.Unmarshal(item.Value, &user)

	db, err := util.GetDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": "can't access database",
		})
		return
	}
	defer db.Close()

	db.Where("owner = ?", user.ID).Find(&files)
	c.JSON(200, files)
}

func CreateFile(c *gin.Context) {
	// todo: 创建文章权限检查
	title := c.PostForm("title")
	content := c.PostForm("content")
	categoryStr := c.PostForm("category")
	if title == "" || content == "" {
		c.JSON(400, gin.H{
			"err": "title or content is empty",
		})
		return
	}
	var category int
	var err error
	if categoryStr != "" {
		category, err = strconv.Atoi(categoryStr)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"err": "category not found",
			})
			return
		}
		// todo: category 验证
		// 暂时先为默认
		category = 0
	} else {
		category = 0
	}

	// 获取登录信息
	token, err := c.Cookie("id")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"err": "not found cookie",
		})
		return
	}
	item, err := util.GetCache().Get(token)
	if item == nil || err != nil {
		c.JSON(500, gin.H{
			"err": "user not found",
		})
		return
	}

	var user module.User
	json.Unmarshal(item.Value, &user)

	file := module.File{
		ID:       0,
		Owner:    user.ID,
		Category: &category,
		Content:  content,
		Title:    title,
		DateTime: time.Now(),
	}

	db, err := util.GetDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": "can't create file",
		})
		return
	}
	defer db.Close()

	db.Create(&file)
	c.JSON(200, gin.H{
		"msg": "create success",
	})
}

func DeleteFile(c *gin.Context) {
	// todo: 删除文章权限检查
	s := c.PostForm("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"err": "file not found",
		})
		return
	}
	db, err := util.GetDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": "can't access database",
		})
		return
	}
	defer db.Close()

	file := module.File{ID: id}
	if err := db.Delete(&file).Error; err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"err": "record not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "delete ok",
	})
}
