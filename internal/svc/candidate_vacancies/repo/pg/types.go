package pg

import (
	"context"
	db "golang-crud/pkg/db/pg"
	"time"
)

type CandidateVacancy struct {
	Id          string    `db:"id"`
	CandidateId string    `db:"candidate_id"`
	VacancyId   string    `db:"vacancy_id"`
	CreatedAt   time.Time `db:"created_at"`
}

type CandidateVacanciesPGRepo struct {
	db db.DBops
}

func NewCandidateVacancyPGRepo(db db.DBops) *CandidateVacanciesPGRepo {
	return &CandidateVacanciesPGRepo{db: db}
}

type CandidateVacancyRepository interface {
	Create(ctx context.Context, dto CandidateVacancy) error
	Delete(ctx context.Context, vacancyId string, candidateId string) error
	GetCandidatesByVacancyId(ctx context.Context, vacancyId string) ([]*string, error)
	GetVacanciesByCandidateId(ctx context.Context, candidateId string) ([]*string, error)
}
