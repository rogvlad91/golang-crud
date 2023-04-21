package memcached

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

const EXP_TIME = 10

func (c CandidateVacancyMemcachedRepo) Add(_ context.Context, items []*string, id string, key string) error {
	marshalledItems, _ := json.Marshal(items)
	err := c.memcached.Set(&memcache.Item{
		Key:        formKeyString(id, key),
		Value:      marshalledItems,
		Expiration: EXP_TIME,
	})

	return err
}

func (c CandidateVacancyMemcachedRepo) Get(_ context.Context, id string, key string) ([]*string, error) {
	result, err := c.memcached.Get(formKeyString(id, key))
	if result != nil {
		var ids []*string
		err = json.Unmarshal(result.Value, &ids)
		return ids, err
	}
	return nil, err
}

func formKeyString(id string, key string) string {
	return fmt.Sprintf("%s:%s", id, key)
}
