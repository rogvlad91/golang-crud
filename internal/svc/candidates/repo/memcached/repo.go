package memcached

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

const EXP_TIME = 10

func (c CandidatesMemcachedRepo) Add(_ context.Context, candidate *Candidate) error {
	marshalledCandidate, _ := json.Marshal(candidate)
	err := c.memcached.Set(&memcache.Item{
		Key:        fmt.Sprintf("CANDIDATE:%s", candidate.Id),
		Value:      marshalledCandidate,
		Expiration: EXP_TIME,
	})
	return err
}

func (c CandidatesMemcachedRepo) Get(_ context.Context, candidateId string) (*Candidate, error) {
	result, err := c.memcached.Get(fmt.Sprintf("CANDIDATE:%s", candidateId))

	if result != nil {
		var candidate *Candidate
		err = json.Unmarshal(result.Value, &candidate)
		return candidate, err
	}

	return nil, err
}
