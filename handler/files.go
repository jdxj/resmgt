package handler

import (
	"github.com/gin-gonic/gin"
)

func GetUserFiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"444": "442",
	})
}
