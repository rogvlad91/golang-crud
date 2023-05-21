package vacancies

import (
	"context"
	"golang-crud/internal/svc/vacancies/repo/memcached"
)

func (s CacheVacancySvc) Create(ctx context.Context, createDTO CreateVacancyDto) (string, error) {
	id, err := s.svc.Create(ctx, createDTO)
	return id, err
}

func (s CacheVacancySvc) GetById(ctx context.Context, id string) (*Vacancy, error) {
	result, err := s.repo.Get(ctx, id)
	if result == nil {
		actualResult, newErr := s.svc.GetById(ctx, id)
		if newErr != nil {
			return nil, newErr
		}

		oneMoreErr := s.repo.Add(ctx, (*memcached.Vacancy)(actualResult))
		if oneMoreErr != nil {
			return nil, oneMoreErr
		}
		return actualResult, nil
	}

	return (*Vacancy)(result), err
}

func (s CacheVacancySvc) GetAll(ctx context.Context) ([]*Vacancy, error) {
	result, err := s.svc.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CacheVacancySvc) Update(ctx context.Context, id string, dto UpdateVacancyDto) error {
	return s.svc.Update(ctx, id, dto)
}

func (s CacheVacancySvc) Delete(ctx context.Context, id string) error {
	return s.svc.Delete(ctx, id)
}
