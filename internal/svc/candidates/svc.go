package candidates

import (
	"context"
	"github.com/gofrs/uuid"
	"golang-crud/internal/svc/candidates/repo/pg"
	"time"
)

func (s CandidateSvc) Create(ctx context.Context, createDTO CreateCandidateDTO) (string, error) {
	id, _ := uuid.NewV4()
	createdAt := time.Now()
	updatedAt := time.Now()
	var candidate = &pg.Candidate{
		Id:           id.String(),
		FullName:     createDTO.FullName,
		Age:          createDTO.Age,
		WantedJob:    createDTO.WantedJob,
		WantedSalary: createDTO.WantedSalary,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	err := s.repo.Create(ctx, *candidate)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s CandidateSvc) GetById(ctx context.Context, id string) (*Candidate, error) {
	result, err := s.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &Candidate{
		Id:           id,
		FullName:     result.FullName,
		Age:          result.Age,
		WantedJob:    result.WantedJob,
		WantedSalary: result.WantedSalary,
		CreatedAt:    result.CreatedAt.String(),
		UpdatedAt:    result.UpdatedAt.String(),
	}, err
}

func (s CandidateSvc) GetAll(ctx context.Context) ([]*Candidate, error) {
	result, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	parsedResult := make([]*Candidate, 0, len(result))

	for _, candidate := range result {
		parsedResult = append(parsedResult, &Candidate{
			Id:           candidate.Id,
			FullName:     candidate.FullName,
			Age:          candidate.Age,
			WantedJob:    candidate.WantedJob,
			WantedSalary: candidate.WantedSalary,
			CreatedAt:    candidate.CreatedAt.String(),
			UpdatedAt:    candidate.UpdatedAt.String(),
		})
	}

	return parsedResult, err
}

func (s CandidateSvc) Update(ctx context.Context, id string, dto UpdateCandidateDto) error {
	updatedAt := time.Now()

	updateDto := &pg.UpdateCandidateDtoExtended{
		WantedJob:    dto.WantedJob,
		WantedSalary: dto.WantedSalary,
		UpdatedAt:    updatedAt,
	}

	err := s.repo.Update(ctx, id, *updateDto)

	return err
}

func (s CandidateSvc) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	return err
}
