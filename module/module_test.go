package module

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"resmgt/util"
	"testing"
	"time"
)

func TestUserMod(t *testing.T) {
	user := User{}
	util.MyDB.First(&user)
	fmt.Println(user)
}

func TestGoMemCached(t *testing.T) {
	mc := memcache.New(":11211")
	item := memcache.Item{
		Key:   "test",
		Value: []byte("Hello World!"),
	}
	mc.Set(&item)

	itemP, err := mc.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(itemP)
}

func TestFileMod(t *testing.T) {
	file := File{}
	util.MyDB.First(&file)
	fmt.Println(file)
}

func TestCategoryMod(t *testing.T) {
	category := Category{}
	util.MyDB.First(&category)
	fmt.Println(category)
}

func TestRoleMod(t *testing.T) {
	role := Role{}
	util.MyDB.First(&role)
	fmt.Println(role)
}

func TestPermissionMod(t *testing.T) {
	perm := Permission{}
	util.MyDB.First(&perm)
	fmt.Println(perm)
}

func TestRolePermMod(t *testing.T) {
	rp := RolePerm{}
	util.MyDB.First(&rp)
	fmt.Println(rp)
}

func TestGormFind(t *testing.T) {
	var files []File
	util.MyDB.Find(&files)
	for _, f := range files {
		fmt.Println(f)
	}
}

func TestFindUser(t *testing.T) {
	var user User
	util.MyDB.Where("name = ? and password = ?", "jdxj", "jdxj").Find(&user)

	fmt.Println(user)
}

func TestInsertFile(t *testing.T) {
	file := File{
		ID:       0,
		Owner:    0,
		Category: nil,
		Content:  "3333",
		Title:    "9001",
		DateTime: time.Now(),
	}
	util.MyDB.Create(&file)
}

func TestDelFile(t *testing.T) {
	file := File{ID: 8}
	if err := util.MyDB.Delete(&file).Error; err != nil {
		panic(err)
	}
}
