//go:generate mockgen -source ./types.go -destination=./mocks/svc.go -package=mock_candidate_vacancy_svc

package candidate_vacancies

import (
	"context"
	"golang-crud/internal/svc/candidate_vacancies/repo/memcached"
	"golang-crud/internal/svc/candidate_vacancies/repo/pg"
)

type CandidateVacancy struct {
	Id          string
	CandidateId string
	VacancyId   string
	CreatedAt   string
}

type CreateDto struct {
	CandidateId string `json:"candidate_id"`
	VacancyId   string `json:"vacancy_id"`
}

type CandidateVacanciesSvc struct {
	repo pg.CandidateVacancyRepository
}

func NewCandidateVacanciesSvc(repo pg.CandidateVacancyRepository) *CandidateVacanciesSvc {
	return &CandidateVacanciesSvc{repo: repo}
}

type CandidateVacanciesCacheSvc struct {
	repo memcached.CandidateVacancyCacheRepository
	svc  CandidateVacanciesProcessor
}

func NewCandidateVacanciesCacheSvc(repo memcached.CandidateVacancyCacheRepository, svc CandidateVacanciesProcessor) *CandidateVacanciesCacheSvc {
	return &CandidateVacanciesCacheSvc{repo: repo, svc: svc}
}

type CandidateVacanciesProcessor interface {
	Create(ctx context.Context, dto CreateDto) (string, error)
	DeleteResponseForVacancy(ctx context.Context, vacancyId string, candidateId string) error
	GetCandidatesByVacancyId(ctx context.Context, vacancyId string) ([]*string, error)
	GetVacanciesByCandidate(ctx context.Context, candidateId string) ([]*string, error)
}
