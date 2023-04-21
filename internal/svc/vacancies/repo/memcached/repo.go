package memcached

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

const EXP_TIME = 10

func (c VacancyMemcachedRepo) Add(_ context.Context, vacancy *Vacancy) error {
	marshalledVacancy, _ := json.Marshal(vacancy)
	err := c.memcached.Set(&memcache.Item{
		Key:        fmt.Sprintf("VACANCY:%s", vacancy.Id),
		Value:      marshalledVacancy,
		Expiration: EXP_TIME,
	})
	return err
}

func (c VacancyMemcachedRepo) Get(_ context.Context, vacancyId string) (*Vacancy, error) {
	result, err := c.memcached.Get(fmt.Sprintf("VACANCY:%s", vacancyId))

	if result != nil {
		var vacancy *Vacancy
		err = json.Unmarshal(result.Value, &vacancy)
		return vacancy, err
	}

	return nil, err
}
