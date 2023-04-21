package memcached

import "github.com/bradfitz/gomemcache/memcache"

const memcachedHost = "localhost:11211"

func NewMemcachedClient() *memcache.Client {
	return memcache.New(memcachedHost)
}
