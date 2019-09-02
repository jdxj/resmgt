package handler

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(200, gin.H{
		"user": user,
	})
}
