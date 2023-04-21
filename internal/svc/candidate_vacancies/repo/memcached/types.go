package memcached

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
)

type CandidateVacancyMemcachedRepo struct {
	memcached *memcache.Client
}

func NewCandidateVacancyMemcachedRepo(memcached *memcache.Client) *CandidateVacancyMemcachedRepo {
	return &CandidateVacancyMemcachedRepo{memcached: memcached}
}

type CandidateVacancyCacheRepository interface {
	Add(ctx context.Context, items []*string, id string, key string) error
	Get(ctx context.Context, id string, key string) ([]*string, error)
}
