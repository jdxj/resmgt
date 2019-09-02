package handler

import (
	"github.com/gin-gonic/gin"
	"resmgt/module"
	"resmgt/util"
)

func Home(c *gin.Context) {
	db, err := util.GetDB()
	if err != nil {
		return
	}
	defer db.Close()

	var files []module.File
	db.Find(&files)

	c.JSON(200, files)
}
