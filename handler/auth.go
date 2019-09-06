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

	var user module.User
	util.MyDB.Where("name = ? and password = ?", name, password).Find(&user)

	if user.Name == "" { // 没找到
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
	err := util.GetCache().Set(&item)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"msg": "login success",
	})
}

// Authenticate 用于验证用户是否登录
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

var permissions = []string{
	"", // 占位
	"创建文章",
	"修改自己的文章",
	"删除自己的文章",
	"浏览自己的文章",
	"浏览别人的文章",
}

const (
	Create = iota + 1
	Modify
	Delete
	BrowseOwn
	BrowOther
	CateCrete
	CateModify
	CateDelete
	CateBrowse
	CateBroOth
)

// AuthGet 验证用户是否有获取其自己 file 的权限
func AuthGet(c *gin.Context) {
	if !checkPerm(c, BrowseOwn) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on get",
		})
	}
}

func AuthCreate(c *gin.Context) {
	if !checkPerm(c, Create) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on create",
		})
	}
}

func AuthDelete(c *gin.Context) {
	if !checkPerm(c, Delete) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on create",
		})
	}
}

// curUser 用于从 memcache 中获取 user 信息
func curUser(c *gin.Context) *module.User {
	token, _ := c.Cookie("id")

	item, err := util.GetCache().Get(token)
	if item == nil || err != nil {
		fmt.Println("err:", err)
		fmt.Println("item:", item)
		return nil
	}

	var user module.User
	json.Unmarshal(item.Value, &user)
	return &user
}

// checkPerm 检查 user 是否具有指定权限
func checkPerm(c *gin.Context, perm int) bool {
	user := curUser(c)
	if user == nil {
		return false
	}
	if user.Role == nil { // 默认对 null 记录的用户采取拒绝
		return false
	}
	if *user.Role == 0 { // 管理员直接过
		return true
	}

	var rolePerms []module.RolePerm
	util.MyDB.Where("role = ?", user.Role).Find(&rolePerms)
	for _, rp := range rolePerms {
		if rp.Perm == perm {
			return true
		}
	}
	return false
}

func AuthCateCreate(c *gin.Context) {
	if !checkPerm(c, CateCrete) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on create cate",
		})
	}
}

func AuthCateModify(c *gin.Context) {
	if !checkPerm(c, CateModify) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on modify cate",
		})
	}

}

func AuthCateDelete(c *gin.Context) {
	if !checkPerm(c, CateDelete) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on delete cate",
		})
	}
}

func AuthCateBrowse(c *gin.Context) {
	if !checkPerm(c, CateBrowse) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on browse cate",
		})
	}
}

func AuthCateBroOth(c *gin.Context) {
	if !checkPerm(c, CateBroOth) {
		c.AbortWithStatusJSON(400, gin.H{
			"err": "permission denied on browse other cate",
		})
	}
}
