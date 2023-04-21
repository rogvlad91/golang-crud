package vacancy

import (
	"context"
	"homework-7/internal/svc/vacancies"
)

type VacanciesController struct {
	svc vacancies.VacancyProcessor
}

func NewVacanciesController(svc vacancies.VacancyProcessor) *VacanciesController {
	return &VacanciesController{svc: svc}
}

type VacanciesCLIProcessor interface {
	Create(ctx context.Context, title string, salary string) error
	GetById(ctx context.Context, id string) error
	GetAll(ctx context.Context) error
	Update(ctx context.Context, id string, title string, salary string) error
	Delete(ctx context.Context, id string) error
}
