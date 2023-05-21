//go:build integration
// +build integration

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-crud/internal/svc/candidates/repo/pg"
	"testing"
	"time"
)

func TestCandidateRepo(t *testing.T) {
	t.Run("createCandidate", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			candidate := pg.Candidate{
				Id:           "1",
				FullName:     "Test candidate",
				WantedSalary: 1000,
				WantedJob:    "President",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			err := repo.Create(ctx, candidate)
			assert.NoError(t, err)

			c, err := repo.GetById(ctx, "1")
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, candidate.FullName, c.FullName)
		})
	})
	t.Run("getCandidateById", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			candidate := pg.Candidate{
				Id:           "1",
				FullName:     "Test candidate",
				WantedSalary: 1000,
				WantedJob:    "President",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			err := repo.Create(ctx, candidate)
			assert.NoError(t, err)

			c, err := repo.GetById(ctx, "1")
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, candidate.FullName, c.FullName)
		})
		t.Run("failure not found candidate", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			c, err := repo.GetById(ctx, "1488")
			assert.ErrorIs(t, err, pg.CandidateNotFoundError)
			assert.Nil(t, c)
		})
	})
	t.Run("getAll", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			candidates := []*pg.Candidate{
				{
					Id:           "1",
					FullName:     "Test candidate",
					WantedSalary: 1000,
					WantedJob:    "President",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				},
				{
					Id:           "2",
					FullName:     "Test candidate 2",
					WantedSalary: 10000,
					WantedJob:    "Minister",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				},
			}
			for _, v := range candidates {
				repo.Create(ctx, *v)
			}

			c, err := repo.GetAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(candidates), len(c))
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			candidate := pg.Candidate{
				Id:           "1",
				FullName:     "Test candidate",
				WantedSalary: 1000,
				WantedJob:    "President",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			repo.Create(ctx, candidate)

			update := pg.UpdateCandidateDtoExtended{
				WantedJob:    "Minister",
				WantedSalary: 2222,
				UpdatedAt:    time.Now(),
			}
			err := repo.Update(ctx, "1", update)
			assert.NoError(t, err)

			c, err := repo.GetById(ctx, "1")
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, c.WantedJob, update.WantedJob)
		})
		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			update := pg.UpdateCandidateDtoExtended{
				WantedJob:    "Minister",
				WantedSalary: 2222,
				UpdatedAt:    time.Now(),
			}
			err := repo.Update(ctx, "1", update)
			assert.ErrorIs(t, err, pg.CandidateNotFoundError)
		})
	})
	t.Run("delete", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			candidate := pg.Candidate{
				Id:           "1",
				FullName:     "Test candidate",
				WantedSalary: 1000,
				WantedJob:    "President",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			repo.Create(ctx, candidate)

			err := repo.Delete(ctx, "1")
			assert.NoError(t, err)
			c, err := repo.GetById(ctx, "1")
			assert.ErrorIs(t, err, pg.CandidateNotFoundError)
			assert.Nil(t, c)
		})
		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewCandidatePGRepo(Database.Db)

			err := repo.Delete(ctx, "1")
			assert.ErrorIs(t, err, pg.CandidateNotFoundError)
		})
	})
}
