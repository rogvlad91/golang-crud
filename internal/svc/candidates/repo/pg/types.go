package pg

import (
	"context"
	db "golang-crud/pkg/db/pg"
	"time"
)

type CandidateModel struct {
	Id           string    `db:"id"`
	FullName     string    `db:"full_name"`
	Age          int       `db:"age"`
	WantedJob    string    `db:"wanted_job"`
	WantedSalary int       `db:"wanted_salary"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Candidate struct {
	Id           string
	FullName     string
	Age          int
	WantedJob    string
	WantedSalary int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UpdateCandidateDtoExtended struct {
	WantedJob    string
	WantedSalary int
	UpdatedAt    time.Time
}

type CandidatePGRepo struct {
	db db.DBops
}

func NewCandidatePGRepo(db db.DBops) *CandidatePGRepo {
	return &CandidatePGRepo{db: db}
}

type CandidateRepository interface {
	Create(ctx context.Context, createDTO Candidate) error
	GetById(ctx context.Context, id string) (*CandidateModel, error)
	GetAll(ctx context.Context) ([]*CandidateModel, error)
	Update(ctx context.Context, id string, dto UpdateCandidateDtoExtended) error
	Delete(ctx context.Context, id string) error
}
