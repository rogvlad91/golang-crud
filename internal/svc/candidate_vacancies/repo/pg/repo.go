package pg

import (
	"context"
	"errors"
)

var ResponseNotFoundError = errors.New("response not found")

func (r CandidateVacanciesPGRepo) Create(ctx context.Context, dto CandidateVacancy) error {
	_, err := r.db.Exec(ctx, CreateQuery, dto.Id, dto.CandidateId, dto.VacancyId, dto.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r CandidateVacanciesPGRepo) Delete(ctx context.Context, vacancyId string, candidateId string) error {
	cTag, err := r.db.Exec(ctx, DeleteQuery, vacancyId, candidateId)
	if err != nil {
		return err
	}
	if cTag.RowsAffected() != 1 {
		return ResponseNotFoundError
	}

	return nil
}

func (r CandidateVacanciesPGRepo) GetCandidatesByVacancyId(ctx context.Context, vacancyId string) ([]*string, error) {
	candidatesResult := make([]*string, 0)
	err := r.db.Select(ctx, &candidatesResult, GetCandidatesByVacancyIdQuery, vacancyId)
	return candidatesResult, err
}

func (r CandidateVacanciesPGRepo) GetVacanciesByCandidateId(ctx context.Context, candidateId string) ([]*string, error) {
	vacancyResult := make([]*string, 0)
	err := r.db.Select(ctx, &vacancyResult, GetVacanciesByCandidateIdQuery, candidateId)
	return vacancyResult, err
}
