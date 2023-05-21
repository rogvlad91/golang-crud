package apps

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"golang-crud/pkg/db/memcached"
	db "golang-crud/pkg/db/pg"
	"log"
)

func BuildPg(ctx context.Context) *db.Database {
	pgRepo, dbErr := db.NewDB(ctx)
	if dbErr != nil {
		log.Fatalf("Error while connecting to pg %s", dbErr)
	}
	return pgRepo.Db
}

func BuildMemcached() *memcache.Client {
	return memcached.NewMemcachedClient()
}
