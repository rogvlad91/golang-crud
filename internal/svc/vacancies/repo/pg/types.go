package pg

import (
	"context"
	db "homework-7/pkg/db/pg"
	"time"
)

type VacancyModel struct {
	Id        string    `db:"id"`
	Title     string    `db:"title"`
	Salary    int       `db:"salary"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type VacancyDTOExtended struct {
	Id        string
	Title     string
	Salary    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateVacancyDtoExtended struct {
	Title     string
	Salary    int
	UpdatedAt time.Time
}

type VacancyPGRepo struct {
	db db.DBops
}

func NewVacancyPGRepo(db db.DBops) *VacancyPGRepo {
	return &VacancyPGRepo{db: db}
}

type VacancyRepository interface {
	Create(ctx context.Context, createDTO VacancyDTOExtended) error
	GetById(ctx context.Context, id string) (*VacancyModel, error)
	GetAll(ctx context.Context) ([]*VacancyModel, error)
	Update(ctx context.Context, id string, dto UpdateVacancyDtoExtended) error
	Delete(ctx context.Context, id string) error
}
