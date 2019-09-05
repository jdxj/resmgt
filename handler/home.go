package handler

import (
	"github.com/gin-gonic/gin"
	"resmgt/module"
	"resmgt/util"
)

func Home(c *gin.Context) {
	var files []module.File
	util.MyDB.Find(&files)

	c.JSON(200, files)
}
