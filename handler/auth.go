package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
	"resmgt/module"
	"resmgt/util"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	// todo: 更安全的验证方式
	name := c.PostForm("name")
	password := c.PostForm("password")

	db, err := util.GetDB()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"err": "data illegal",
		})
		return
	}
	defer db.Close()

	var user module.User
	db.Where("name = ? and password = ?", name, password).Find(&user)

	if user.Name == "" { // 没找到
		fmt.Println(err)
		c.AbortWithStatusJSON(400, gin.H{
			"err": "user name or password error",
		})
		return
	}

	// 设置 cookie
	key := "id"
	timestamp := strconv.Itoa(user.ID) + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	maxAge := 3600
	// todo: 域名可能要换
	domain := "localhost"
	c.SetCookie(key, timestamp, maxAge, "/", domain, false, true)

	// 设置 memcache
	value, _ := json.Marshal(&user)
	item := memcache.Item{
		Key:   timestamp,
		Value: value,
	}
	err = util.GetCache().Set(&item)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"err": "error to cache",
		})
		return
	}
}

func Authenticate(c *gin.Context) {
	token, err := c.Cookie("id")
	if err != nil { // 没有找到?
		fmt.Println(err)
		c.AbortWithStatusJSON(400, gin.H{
			"err": "not login",
		})
		return
	}

	item, err := util.GetCache().Get(token)
	if item == nil || err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "not login 2",
		})
		return
	}
}
