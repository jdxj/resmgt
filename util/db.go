package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var MyDB = getDB()

func getDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/resmgt?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.DB().SetConnMaxLifetime(time.Hour)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetMaxIdleConns(2)

	return db
}
