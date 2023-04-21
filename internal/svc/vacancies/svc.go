package vacancies

import (
	"context"
	"github.com/gofrs/uuid"
	"homework-7/internal/svc/vacancies/repo/pg"
	"time"
)

func (s VacancySvc) Create(ctx context.Context, createDTO CreateVacancyDto) (string, error) {
	id, _ := uuid.NewV4()
	createdAt := time.Now()
	updatedAt := time.Now()
	var vacancy = &VacancyDTOExtended{
		id.String(),
		createDTO.Title,
		createDTO.Salary,
		createdAt,
		updatedAt,
	}

	err := s.repo.Create(ctx, pg.VacancyDTOExtended(*vacancy))
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s VacancySvc) GetById(ctx context.Context, id string) (*Vacancy, error) {
	result, err := s.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &Vacancy{
		Id:        result.Id,
		Title:     result.Title,
		Salary:    result.Salary,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	}, err
}

func (s VacancySvc) GetAll(ctx context.Context) ([]*Vacancy, error) {
	result, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	parsedResult := make([]*Vacancy, 0, len(result))

	for _, vacancy := range result {
		parsedResult = append(parsedResult, &Vacancy{
			Id:        vacancy.Id,
			Title:     vacancy.Title,
			Salary:    vacancy.Salary,
			CreatedAt: vacancy.CreatedAt.String(),
			UpdatedAt: vacancy.UpdatedAt.String(),
		})
	}

	return parsedResult, err
}

func (s VacancySvc) Update(ctx context.Context, id string, dto UpdateVacancyDto) error {
	updatedAt := time.Now()

	updateDto := &pg.UpdateVacancyDtoExtended{
		Title:     dto.Title,
		Salary:    dto.Salary,
		UpdatedAt: updatedAt,
	}

	err := s.repo.Update(ctx, id, *updateDto)

	return err
}

func (s VacancySvc) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	return err
}
