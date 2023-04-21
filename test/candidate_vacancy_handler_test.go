//go:build integration
// +build integration

package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"homework-7/apps"
	candidate_vacancies_controller "homework-7/apps/http/candidate_vacancies"
	"homework-7/internal/svc/candidate_vacancies"
	memcached2 "homework-7/internal/svc/candidate_vacancies/repo/memcached"
	"homework-7/internal/svc/candidate_vacancies/repo/pg"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCandidateVacancyHandler(t *testing.T) {
	t.Run("Create Response", func(t *testing.T) {
		Database.SetUp(t)
		defer Database.TearDown(t)
		memcached := apps.BuildMemcached()

		repo := pg.NewCandidateVacancyPGRepo(Database.Db)
		cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
		svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
		cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

		controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
		router := candidate_vacancies_controller.GetRouter(controller)
		server := httptest.NewServer(router)
		defer server.Close()

		testCases := []struct {
			name        string
			inputJSON   string
			expectedErr bool
		}{
			{
				name:        "Valid input",
				inputJSON:   `{"candidate_id": "123", "vacancy_id": "456"}`,
				expectedErr: false,
			},
			{
				name:        "Missing candidate ID",
				inputJSON:   `{"vacancy_id": "456"}`,
				expectedErr: true,
			},
			{
				name:        "Missing vacancy ID",
				inputJSON:   `{"candidate_id": "123"}`,
				expectedErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				reqBody := strings.NewReader(tc.inputJSON)
				req, err := http.NewRequest("POST", server.URL+"/", reqBody)
				assert.NoError(t, err)

				resp, err := http.DefaultClient.Do(req)
				assert.NoError(t, err)
				defer resp.Body.Close()

				if tc.expectedErr {
					assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				} else {
					assert.Equal(t, http.StatusCreated, resp.StatusCode)
				}

				if !tc.expectedErr {
					var responseBody struct {
						Id string `json:"id"`
					}
					err = json.NewDecoder(resp.Body).Decode(&responseBody)
					assert.NoError(t, err)

					assert.NotEmpty(t, responseBody.Id)
				}
			})
		}
	})
	t.Run("DeleteResponse", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
			svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
			cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

			controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
			router := candidate_vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()
			inputJSON := `{"candidate_id": "123", "vacancy_id": "456"}`
			reqBody := strings.NewReader(inputJSON)
			req, err := http.NewRequest("POST", server.URL+"/", reqBody)
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			deleteReq, err := http.NewRequest("DELETE", fmt.Sprintf("%s?vacancy_id=456&candidate_id=123", server.URL), nil)
			deleteResp, err := http.DefaultClient.Do(deleteReq)
			assert.NoError(t, err)
			defer deleteResp.Body.Close()

			getReq, err := http.NewRequest("GET", fmt.Sprintf("%s/candidates?vacancy_id=456", server.URL), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			assert.NoError(t, err)
			defer getResp.Body.Close()

			assert.Equal(t, http.StatusNotFound, getResp.StatusCode)

		})

		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
			svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
			cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

			controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
			router := candidate_vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			deleteReq, err := http.NewRequest("DELETE", fmt.Sprintf("%s?vacancy_id=456&candidate_id=123", server.URL), nil)
			deleteResp, err := http.DefaultClient.Do(deleteReq)
			assert.NoError(t, err)
			defer deleteResp.Body.Close()
			assert.Equal(t, http.StatusNotFound, deleteResp.StatusCode)
		})
	})

	t.Run("getCandidatesByVacancy", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
			svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
			cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

			controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
			router := candidate_vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			inputJSON := `{"candidate_id": "456", "vacancy_id": "123"}`
			reqBody := strings.NewReader(inputJSON)
			req, err := http.NewRequest("POST", server.URL+"/", reqBody)
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			getReq, err := http.NewRequest("GET", fmt.Sprintf("%s/candidates?vacancy_id=123", server.URL), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			assert.NoError(t, err)
			defer getResp.Body.Close()

			var candidates []string
			err = json.NewDecoder(getResp.Body).Decode(&candidates)
			assert.NoError(t, err)

			assert.Equal(t, []string{"456"}, candidates)

		})

		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
			svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
			cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

			controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
			router := candidate_vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()
			deleteReq, err := http.NewRequest("GET", fmt.Sprintf("%s/candidates?vacancy_id=14", server.URL), nil)
			deleteResp, err := http.DefaultClient.Do(deleteReq)
			assert.NoError(t, err)
			defer deleteResp.Body.Close()
			assert.Equal(t, http.StatusNotFound, deleteResp.StatusCode)
		})
	})

	t.Run("getVacanciesByVacancy", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
			svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
			cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

			controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
			router := candidate_vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			inputJSON := `{"candidate_id": "123", "vacancy_id": "456"}`
			reqBody := strings.NewReader(inputJSON)
			req, err := http.NewRequest("POST", server.URL+"/", reqBody)
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			getReq, err := http.NewRequest("GET", fmt.Sprintf("%s/vacancies?candidate_id=123", server.URL), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			assert.NoError(t, err)
			defer getResp.Body.Close()

			var vacancies []string
			err = json.NewDecoder(getResp.Body).Decode(&vacancies)
			assert.NoError(t, err)

			assert.Equal(t, []string{"456"}, vacancies)

		})

		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewCandidateVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewCandidateVacancyMemcachedRepo(memcached)
			svc := candidate_vacancies.NewCandidateVacanciesSvc(repo)
			cacheSvc := candidate_vacancies.NewCandidateVacanciesCacheSvc(cacheRepo, svc)

			controller := candidate_vacancies_controller.NewCandidateVacanciesController(cacheSvc)
			router := candidate_vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			deleteReq, err := http.NewRequest("GET", fmt.Sprintf("%s/vacancies?candidate_id=14", server.URL), nil)
			deleteResp, err := http.DefaultClient.Do(deleteReq)
			assert.NoError(t, err)
			defer deleteResp.Body.Close()
			assert.Equal(t, http.StatusNotFound, deleteResp.StatusCode)
		})
	})
}
