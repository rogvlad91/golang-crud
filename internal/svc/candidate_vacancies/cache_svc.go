package candidate_vacancies

import (
	"context"
)

const (
	CANDIDATE_KEY = "CANDIDATES"
	VACANCY_KEY   = "VACANCIES"
)

func (s CandidateVacanciesCacheSvc) Create(ctx context.Context, dto CreateDto) (string, error) {
	return s.svc.Create(ctx, dto)
}

func (s CandidateVacanciesCacheSvc) DeleteResponseForVacancy(ctx context.Context, vacancyId string, candidateId string) error {
	err := s.svc.DeleteResponseForVacancy(ctx, vacancyId, candidateId)
	return err
}

func (s CandidateVacanciesCacheSvc) GetCandidatesByVacancyId(ctx context.Context, vacancyId string) ([]*string, error) {
	result, err := s.repo.Get(ctx, vacancyId, CANDIDATE_KEY)

	if result == nil {
		actualResult, newErr := s.svc.GetCandidatesByVacancyId(ctx, vacancyId)
		if newErr != nil {
			return nil, newErr
		}
		oneMoreError := s.repo.Add(ctx, actualResult, vacancyId, CANDIDATE_KEY)
		if oneMoreError != nil {
			return nil, oneMoreError
		}
		return actualResult, nil
	}

	return result, err
}

func (s CandidateVacanciesCacheSvc) GetVacanciesByCandidate(ctx context.Context, candidateId string) ([]*string, error) {
	result, err := s.repo.Get(ctx, candidateId, VACANCY_KEY)

	if result == nil {
		actualResult, newErr := s.svc.GetVacanciesByCandidate(ctx, candidateId)
		if newErr != nil {
			return nil, newErr
		}
		oneMoreError := s.repo.Add(ctx, actualResult, candidateId, VACANCY_KEY)
		if oneMoreError != nil {
			return nil, oneMoreError
		}
		return actualResult, nil

	}

	return result, err
}
