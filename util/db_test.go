package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

// TestConnCount 测试连接数上限
func TestConnCount(*testing.T) {
	var dbs []*gorm.DB
	count := 0
	for {
		db := getDB()
		dbs = append(dbs, db)
		count++
		fmt.Println("count:", count)
		time.Sleep(500 * time.Millisecond)
	}
}

// TestConnPool 测试连接池
func TestConnPool(t *testing.T) {
	err := MyDB.DB().Ping()
	if err != nil {
		panic(err)
	}

	MyDB.DB().SetConnMaxLifetime(time.Hour)
	MyDB.DB().SetMaxIdleConns(1)
	MyDB.DB().SetMaxOpenConns(1)
}
