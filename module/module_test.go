package module

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"resmgt/util"
	"testing"
)

func TestUserMod(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := User{}
	db.First(&user)
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
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	file := File{}
	db.First(&file)
	fmt.Println(file)
}

func TestCategoryMod(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	category := Category{}
	db.First(&category)
	fmt.Println(category)
}

func TestRoleMod(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	role := Role{}
	db.First(&role)
	fmt.Println(role)
}

func TestPermissionMod(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	perm := Permission{}
	db.First(&perm)
	fmt.Println(perm)
}

func TestRolePermMod(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rp := RolePerm{}
	db.First(&rp)
	fmt.Println(rp)
}

func TestGormFind(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var files []File
	db.Find(&files)
	for _, f := range files {
		fmt.Println(f)
	}
}

func TestFindUser(t *testing.T) {
	db, err := util.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.Where("name = ? and password = ?", "jdxj", "jdxj").Find(&user)

	fmt.Println(user)
}
