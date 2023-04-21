//go:build integration
// +build integration

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"homework-7/internal/svc/candidate_vacancies/repo/pg"
	"testing"
	"time"
)

func TestCandidateVacancyRepo(t *testing.T) {
	t.Run("createResponse", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)

			candidateVacancy := pg.CandidateVacancy{
				Id:          "cv-1",
				CandidateId: "candidate-1",
				VacancyId:   "vacancy-1",
				CreatedAt:   time.Now(),
			}

			err := repo.Create(ctx, candidateVacancy)
			assert.NoError(t, err)

			candidates, err := repo.GetCandidatesByVacancyId(ctx, "vacancy-1")
			assert.NoError(t, err)
			assert.Equal(t, len(candidates), 1)
			assert.Equal(t, "candidate-1", *candidates[0])
		})
	})
	t.Run("DeleteResponse", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)

			candidateVacancy := pg.CandidateVacancy{
				Id:          "cv-1",
				CandidateId: "candidate-1",
				VacancyId:   "vacancy-1",
				CreatedAt:   time.Now(),
			}

			err := repo.Create(ctx, candidateVacancy)
			assert.NoError(t, err)

			err = repo.Delete(ctx, candidateVacancy.VacancyId, candidateVacancy.CandidateId)
			assert.NoError(t, err)

			candidates, err := repo.GetCandidatesByVacancyId(ctx, "vacancy-1")
			assert.Empty(t, candidates)
		})
		t.Run("failure not found candidate", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)

			err := repo.Delete(ctx, "1", "2")

			assert.ErrorIs(t, err, pg.ResponseNotFoundError)
		})
	})
	t.Run("getCandidatesByVacancy", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)

			candidateVacancy := pg.CandidateVacancy{
				Id:          "cv-1",
				CandidateId: "candidate-1",
				VacancyId:   "vacancy-1",
				CreatedAt:   time.Now(),
			}

			err := repo.Create(ctx, candidateVacancy)
			assert.NoError(t, err)

			candidates, err := repo.GetCandidatesByVacancyId(ctx, "vacancy-1")
			assert.NoError(t, err)
			assert.Equal(t, len(candidates), 1)
			assert.Equal(t, "candidate-1", *candidates[0])
		})
	})

	t.Run("GetVacanciesByCandidate", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)

			candidateVacancy := pg.CandidateVacancy{
				Id:          "cv-1",
				CandidateId: "candidate-1",
				VacancyId:   "vacancy-1",
				CreatedAt:   time.Now(),
			}

			err := repo.Create(ctx, candidateVacancy)
			assert.NoError(t, err)

			vacancies, err := repo.GetVacanciesByCandidateId(ctx, "candidate-1")
			assert.NoError(t, err)
			assert.Equal(t, len(vacancies), 1)
			assert.Equal(t, "vacancy-1", *vacancies[0])
		})
	})
}
