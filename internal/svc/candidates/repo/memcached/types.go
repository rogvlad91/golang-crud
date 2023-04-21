package memcached

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
)

type Candidate struct {
	Id           string
	FullName     string
	Age          int
	WantedJob    string
	WantedSalary int
	CreatedAt    string
	UpdatedAt    string
}

type CandidatesMemcachedRepo struct {
	memcached *memcache.Client
}

func NewCandidatesMemcachedRepo(memcached *memcache.Client) *CandidatesMemcachedRepo {
	return &CandidatesMemcachedRepo{memcached: memcached}
}

type CandidatesCacheRepository interface {
	Add(ctx context.Context, candidate *Candidate) error
	Get(ctx context.Context, candidateId string) (*Candidate, error)
}
