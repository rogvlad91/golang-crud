package memcached

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
)

type Vacancy struct {
	Id        string
	Title     string
	Salary    int
	CreatedAt string
	UpdatedAt string
}

type VacancyMemcachedRepo struct {
	memcached *memcache.Client
}

func NewVacancyMemcachedRepo(memcached *memcache.Client) *VacancyMemcachedRepo {
	return &VacancyMemcachedRepo{memcached: memcached}
}

type VacancyCacheRepository interface {
	Add(ctx context.Context, vacancy *Vacancy) error
	Get(ctx context.Context, vacancyId string) (*Vacancy, error)
}
