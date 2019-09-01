package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:root@/resmgt?charset=utf8mb4&parseTime=True&loc=Local")
}
