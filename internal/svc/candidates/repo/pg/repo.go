//go:generate mockgen -source ./repo.go -destination=./mocks/repository.go -package=mock_repository

package pg

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

var CandidateNotFoundError = errors.New("candidate not found")

func (r CandidatePGRepo) Create(ctx context.Context, createDTO Candidate) error {
	cTag, err := r.db.Exec(ctx, CreateQuery, createDTO.Id, createDTO.FullName, createDTO.Age, createDTO.WantedJob, createDTO.WantedSalary, createDTO.CreatedAt, createDTO.UpdatedAt)
	if err != nil {
		return err
	}
	if cTag.RowsAffected() != 1 {
		return errors.New("error creating candidate")
	}

	return nil
}

func (r CandidatePGRepo) GetById(ctx context.Context, id string) (*CandidateModel, error) {
	var c CandidateModel
	err := r.db.ExecQueryRow(ctx, GetByIdQuery, id).Scan(&c.Id, &c.FullName, &c.Age, &c.WantedJob, &c.WantedSalary, &c.CreatedAt, &c.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, CandidateNotFoundError
		}
		return nil, err
	}

	return &c, nil
}

func (r CandidatePGRepo) GetAll(ctx context.Context) ([]*CandidateModel, error) {
	candidatesResult := make([]*CandidateModel, 0)
	err := r.db.Select(ctx, &candidatesResult, GetAllQuery)
	return candidatesResult, err
}

func (r CandidatePGRepo) Update(ctx context.Context, id string, dto UpdateCandidateDtoExtended) error {
	cTag, err := r.db.Exec(ctx, UpdateQuery, dto.WantedJob, dto.WantedSalary, dto.UpdatedAt, id)

	if cTag.RowsAffected() != 1 {
		return CandidateNotFoundError
	}

	if err != nil {
		return err
	}

	return nil
}

func (r CandidatePGRepo) Delete(ctx context.Context, id string) error {
	cTag, err := r.db.Exec(ctx, DeleteQuery, id)

	if cTag.RowsAffected() != 1 {
		return CandidateNotFoundError
	}

	if err != nil {
		return err
	}

	return nil
}
