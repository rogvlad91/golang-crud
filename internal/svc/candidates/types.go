//go:generate mockgen -source ./types.go -destination=./mocks/svc.go -package=mock_candidate_svc

package candidates

import (
	"context"
	"golang-crud/internal/svc/candidates/repo/memcached"
	"golang-crud/internal/svc/candidates/repo/pg"
)

type Candidate struct {
	Id           string
	FullName     string
	Age          int
	WantedJob    string
	WantedSalary int
	CreatedAt    string
	UpdatedAt    string
}

type CreateCandidateDTO struct {
	FullName     string `json:"fullName"`
	Age          int    `json:"age"`
	WantedJob    string `json:"wantedJob"`
	WantedSalary int    `json:"wantedSalary"`
}

type UpdateCandidateDto struct {
	WantedJob    string `json:"wantedJob"`
	WantedSalary int    `json:"wantedSalary"`
}

type CandidateSvc struct {
	repo pg.CandidateRepository
}

func NewCandidateSvc(repo pg.CandidateRepository) *CandidateSvc {
	return &CandidateSvc{repo: repo}
}

type CacheCandidateSvc struct {
	repo memcached.CandidatesCacheRepository
	svc  CandidateProcessor
}

func NewCacheCandidateSvc(repo memcached.CandidatesCacheRepository, svc CandidateProcessor) *CacheCandidateSvc {
	return &CacheCandidateSvc{repo: repo, svc: svc}
}

type CandidateProcessor interface {
	Create(ctx context.Context, createDTO CreateCandidateDTO) (string, error)
	GetById(ctx context.Context, id string) (*Candidate, error)
	GetAll(ctx context.Context) ([]*Candidate, error)
	Update(ctx context.Context, id string, dto UpdateCandidateDto) error
	Delete(ctx context.Context, id string) error
}
