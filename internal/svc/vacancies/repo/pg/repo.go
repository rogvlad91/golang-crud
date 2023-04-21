package pg

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

var VacancyNotFoundError = errors.New("vacancy not found")

func (r VacancyPGRepo) Create(ctx context.Context, createDTO VacancyDTOExtended) error {
	_, err := r.db.Exec(ctx, CreateQuery, createDTO.Id, createDTO.Title, createDTO.Salary, createDTO.CreatedAt, createDTO.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r VacancyPGRepo) GetById(ctx context.Context, id string) (*VacancyModel, error) {
	var c VacancyModel
	err := r.db.ExecQueryRow(ctx, GetByIdQuery, id).Scan(&c.Id, &c.Title, &c.Salary, &c.CreatedAt, &c.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, VacancyNotFoundError
		}
		return nil, err
	}

	return &c, nil
}

func (r VacancyPGRepo) GetAll(ctx context.Context) ([]*VacancyModel, error) {
	candidatesResult := make([]*VacancyModel, 0)
	err := r.db.Select(ctx, &candidatesResult, GetAllQuery)
	return candidatesResult, err
}

func (r VacancyPGRepo) Update(ctx context.Context, id string, dto UpdateVacancyDtoExtended) error {
	cTag, err := r.db.Exec(ctx, UpdateQuery, dto.Title, dto.Salary, dto.UpdatedAt, id)

	if cTag.RowsAffected() != 1 {
		return VacancyNotFoundError
	}

	if err != nil {
		return err
	}

	return nil
}

func (r VacancyPGRepo) Delete(ctx context.Context, id string) error {
	cTag, err := r.db.Exec(ctx, DeleteQuery, id)
	if cTag.RowsAffected() != 1 {
		return VacancyNotFoundError
	}
	if err != nil {
		return err
	}

	return nil
}
