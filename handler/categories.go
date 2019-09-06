package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"resmgt/module"
	"resmgt/util"
	"strconv"
)

// GetUserCategories 用于获取用户的分类信息
func GetUserCategories(c *gin.Context) {
	user := curUser(c)
	if user == nil {
		c.JSON(400, gin.H{
			"err": "user not found",
		})
		return
	}

	var cats []module.Category
	util.MyDB.Where("owner = ?", user.ID).Find(&cats)
	c.JSON(200, cats)
}

func CreateCategory(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(400, gin.H{
			"err": "name is empty",
		})
		return
	}
	// 查询是否有同名分类
	user := curUser(c)
	if user == nil {
		c.JSON(400, gin.H{
			"err": "user not found",
		})
		return
	}

	var cats []module.Category
	util.MyDB.Where("owner = ?", user.ID).Find(&cats)
	for _, cat := range cats {
		if cat.Name == name {
			c.JSON(400, gin.H{
				"err": "category name duplicate",
			})
			return
		}
	}

	// 检查父分类
	s := c.PostForm("pid")
	pid, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"err": "pid illegal",
		})
		return
	}

	for _, cat := range cats {
		if cat.ID != pid {
			continue
		}
		// 找到父 id
		category := module.Category{
			ID:    0,
			Name:  name,
			Pid:   &pid,
			Owner: user.ID,
		}
		// 插入数据
		util.MyDB.Create(&category)
		c.JSON(200, gin.H{
			"msg": "create category success",
		})
		return
	}

	c.JSON(400, gin.H{
		"err": "pid not found",
	})
}

func DeleteCategory(c *gin.Context) {
	s := c.PostForm("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"err": "id illegal",
		})
		return
	}

	user := curUser(c)
	if user == nil {
		c.JSON(400, gin.H{
			"err": "user not found",
		})
		return
	}

	category := module.Category{ID: id}
	util.MyDB.Where("owner = ?", user.ID).Delete(&category)
	c.JSON(200, gin.H{
		"msg": "delete category success",
	})
}
