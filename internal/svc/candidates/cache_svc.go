package candidates

import (
	"context"
	"golang-crud/internal/svc/candidates/repo/memcached"
)

func (s CacheCandidateSvc) Create(ctx context.Context, createDTO CreateCandidateDTO) (string, error) {
	id, err := s.svc.Create(ctx, createDTO)
	return id, err
}

func (s CacheCandidateSvc) GetById(ctx context.Context, id string) (*Candidate, error) {
	result, err := s.repo.Get(ctx, id)

	if result == nil {
		actualResult, newErr := s.svc.GetById(ctx, id)
		if newErr != nil {
			return nil, newErr
		}
		oneMoreErr := s.repo.Add(ctx, (*memcached.Candidate)(actualResult))
		if oneMoreErr != nil {
			return nil, oneMoreErr
		}
		return actualResult, nil
	}

	return (*Candidate)(result), err
}

func (s CacheCandidateSvc) GetAll(ctx context.Context) ([]*Candidate, error) {
	result, err := s.svc.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CacheCandidateSvc) Update(ctx context.Context, id string, dto UpdateCandidateDto) error {
	return s.svc.Update(ctx, id, dto)
}

func (s CacheCandidateSvc) Delete(ctx context.Context, id string) error {
	return s.svc.Delete(ctx, id)
}
