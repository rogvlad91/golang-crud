package candidate_vacancies

import (
	"context"
	"github.com/gofrs/uuid"
	"golang-crud/internal/svc/candidate_vacancies/repo/pg"
	"time"
)

func (s CandidateVacanciesSvc) Create(ctx context.Context, dto CreateDto) (string, error) {
	id, _ := uuid.NewV4()
	createdAt := time.Now()

	var candidateVacancy = &pg.CandidateVacancy{
		Id:          id.String(),
		CandidateId: dto.CandidateId,
		VacancyId:   dto.VacancyId,
		CreatedAt:   createdAt,
	}

	err := s.repo.Create(ctx, *candidateVacancy)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s CandidateVacanciesSvc) DeleteResponseForVacancy(ctx context.Context, vacancyId string, candidateId string) error {
	err := s.repo.Delete(ctx, vacancyId, candidateId)
	return err
}

func (s CandidateVacanciesSvc) GetCandidatesByVacancyId(ctx context.Context, vacancyId string) ([]*string, error) {
	result, err := s.repo.GetCandidatesByVacancyId(ctx, vacancyId)

	return result, err
}

func (s CandidateVacanciesSvc) GetVacanciesByCandidate(ctx context.Context, candidateId string) ([]*string, error) {
	result, err := s.repo.GetVacanciesByCandidateId(ctx, candidateId)

	return result, err
}
