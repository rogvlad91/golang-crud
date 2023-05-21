//go:build integration
// +build integration

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-crud/internal/svc/vacancies/repo/pg"
	"testing"
	"time"
)

func TestRepo(t *testing.T) {
	t.Run("createVacancy", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			vacancy := pg.VacancyDTOExtended{
				Id:        "1",
				Title:     "Test vacancy",
				Salary:    1000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			err := repo.Create(ctx, vacancy)
			assert.NoError(t, err)

			c, err := repo.GetById(ctx, "1")
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, vacancy.Title, c.Title)
		})
	})
	t.Run("getVacancyById", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			vacancy := pg.VacancyDTOExtended{
				Id:        "1",
				Title:     "Test vacancy",
				Salary:    1000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			err := repo.Create(ctx, vacancy)
			assert.NoError(t, err)

			c, err := repo.GetById(ctx, "1")
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, vacancy.Title, c.Title)
		})
		t.Run("failure not found vacancy", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			c, err := repo.GetById(ctx, "1488")
			assert.ErrorIs(t, err, pg.VacancyNotFoundError)
			assert.Nil(t, c)
		})
	})
	t.Run("getAll", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			vacancies := []*pg.VacancyDTOExtended{
				{
					Id:        "1",
					Title:     "Test vacancy 1",
					Salary:    1000,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Id:        "2",
					Title:     "Test vacancy 2",
					Salary:    2000,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			}
			for _, v := range vacancies {
				repo.Create(ctx, *v)
			}

			c, err := repo.GetAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(vacancies), len(c))
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			vacancy := pg.VacancyDTOExtended{
				Id:        "1",
				Title:     "Test vacancy 1",
				Salary:    1000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			repo.Create(ctx, vacancy)

			update := pg.UpdateVacancyDtoExtended{
				Title:     "New title",
				Salary:    2000,
				UpdatedAt: time.Now(),
			}
			err := repo.Update(ctx, "1", update)
			assert.NoError(t, err)

			c, err := repo.GetById(ctx, "1")
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, c.Title, update.Title)
		})
		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			update := pg.UpdateVacancyDtoExtended{
				Title:     "New title",
				Salary:    2000,
				UpdatedAt: time.Now(),
			}
			err := repo.Update(ctx, "1", update)
			assert.ErrorIs(t, err, pg.VacancyNotFoundError)
		})
	})
	t.Run("delete", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			vacancy := pg.VacancyDTOExtended{
				Id:        "1",
				Title:     "Test vacancy 1",
				Salary:    1000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			repo.Create(ctx, vacancy)

			err := repo.Delete(ctx, "1")
			assert.NoError(t, err)
			c, err := repo.GetById(ctx, "1")
			assert.ErrorIs(t, err, pg.VacancyNotFoundError)
			assert.Nil(t, c)
		})
		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			ctx := context.Background()
			repo := pg.NewVacancyPGRepo(Database.Db)

			err := repo.Delete(ctx, "1")
			assert.ErrorIs(t, err, pg.VacancyNotFoundError)
		})
	})
}
