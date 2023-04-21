//go:generate mockgen -source ./types.go -destination=./mocks/svc.go -package=mock_vacancy_svc

package vacancies

import (
	"context"
	"homework-7/internal/svc/vacancies/repo/memcached"
	"homework-7/internal/svc/vacancies/repo/pg"
	"time"
)

type Vacancy struct {
	Id        string
	Title     string
	Salary    int
	CreatedAt string
	UpdatedAt string
}

type CreateVacancyDto struct {
	Title  string `json:"title"`
	Salary int    `json:"salary"`
}

type VacancyDTOExtended struct {
	Id        string
	Title     string
	Salary    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateVacancyDto struct {
	Title  string `json:"title"`
	Salary int    `json:"salary"`
}

type VacancySvc struct {
	repo pg.VacancyRepository
}

func NewVacancySvc(repo pg.VacancyRepository) *VacancySvc {
	return &VacancySvc{repo: repo}
}

type CacheVacancySvc struct {
	repo memcached.VacancyCacheRepository
	svc  VacancyProcessor
}

func NewCacheVacancySvc(repo memcached.VacancyCacheRepository, svc VacancyProcessor) *CacheVacancySvc {
	return &CacheVacancySvc{repo: repo, svc: svc}
}

type VacancyProcessor interface {
	Create(ctx context.Context, createDTO CreateVacancyDto) (string, error)
	GetById(ctx context.Context, id string) (*Vacancy, error)
	GetAll(ctx context.Context) ([]*Vacancy, error)
	Update(ctx context.Context, id string, dto UpdateVacancyDto) error
	Delete(ctx context.Context, id string) error
}
