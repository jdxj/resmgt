package util

import "github.com/bradfitz/gomemcache/memcache"

var mc = memcache.New(":11211")

func GetCache() *memcache.Client {
	return mc
}
