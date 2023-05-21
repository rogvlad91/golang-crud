//go:build integration
// +build integration

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-crud/apps"
	vacancies_controller "golang-crud/apps/http/vacancies"
	"golang-crud/internal/svc/vacancies"
	memcached2 "golang-crud/internal/svc/vacancies/repo/memcached"
	"golang-crud/internal/svc/vacancies/repo/pg"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVacancyHandler(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			createDto := vacancies.CreateVacancyDto{
				Title:  "Software Engineer",
				Salary: 100000,
			}

			jsonData, _ := json.Marshal(createDto)
			req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewBuffer(jsonData))
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			var respBody struct{ Id string }
			err = json.NewDecoder(resp.Body).Decode(&respBody)

			assert.NoError(t, err)

			assert.NotEmpty(t, respBody.Id)

		})

		t.Run("fail, invalid body", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			jsonData := []byte(`{"title": 123}`)
			req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewBuffer(jsonData))
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
	})

	t.Run("getById", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			createDto := vacancies.CreateVacancyDto{
				Title:  "Software Engineer",
				Salary: 100000,
			}

			jsonData, _ := json.Marshal(createDto)
			req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewBuffer(jsonData))
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			var respBody struct{ Id string }
			err = json.NewDecoder(resp.Body).Decode(&respBody)

			assert.NoError(t, err)

			assert.NotEmpty(t, respBody.Id)

			getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			defer getResp.Body.Close()
			assert.NoError(t, err)

			assert.NotNil(t, getResp.Body)

			assert.Equal(t, http.StatusOK, getResp.StatusCode)
			var retrievedVacancy *vacancies.Vacancy
			err = json.NewDecoder(getResp.Body).Decode(&retrievedVacancy)
			assert.NoError(t, err)

			assert.Equal(t, createDto.Title, retrievedVacancy.Title)
			assert.Equal(t, createDto.Salary, retrievedVacancy.Salary)
		})
		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/babaaba", server.URL), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			defer getResp.Body.Close()
			assert.NoError(t, err)
			assert.Equal(t, http.StatusNotFound, getResp.StatusCode)
		})
	})

	t.Run("getAll", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			for i := 0; i < 3; i++ {
				createDTO := vacancies.CreateVacancyDto{
					Title:  "President",
					Salary: i * 100000,
				}
				vacancyBytes, _ := json.Marshal(createDTO)
				req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(vacancyBytes))
				resp, err := http.DefaultClient.Do(req)
				assert.NoError(t, err)

				defer resp.Body.Close()
				var respBody struct{ Id string }
				err = json.NewDecoder(resp.Body).Decode(&respBody)
				assert.NoError(t, err)
			}
			req, _ := http.NewRequest("GET", server.URL+"/", nil)
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			defer resp.Body.Close()
			var vacancies []*vacancies.Vacancy
			err = json.NewDecoder(resp.Body).Decode(&vacancies)
			assert.NoError(t, err)

			assert.Equal(t, 3, len(vacancies))

			for i := 0; i < 3; i++ {
				assert.Equal(t, "President", vacancies[i].Title)
				assert.Equal(t, i*100000, vacancies[i].Salary)
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			createDto := vacancies.CreateVacancyDto{
				Title:  "Software Engineer",
				Salary: 100000,
			}

			jsonData, _ := json.Marshal(createDto)
			req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(jsonData))
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			var respBody struct{ Id string }
			err = json.NewDecoder(resp.Body).Decode(&respBody)
			assert.NoError(t, err)

			updateDto := vacancies.UpdateVacancyDto{
				Title:  "President",
				Salary: 2000000,
			}
			jsonUpdateDto, _ := json.Marshal(updateDto)

			updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/%s", server.URL, respBody.Id), bytes.NewBuffer(jsonUpdateDto))
			updateResp, err := http.DefaultClient.Do(updateReq)
			assert.NoError(t, err)

			defer updateResp.Body.Close()
			assert.Equal(t, http.StatusNoContent, updateResp.StatusCode)

			getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			defer getResp.Body.Close()
			assert.NoError(t, err)

			assert.NotNil(t, getResp.Body)

			assert.Equal(t, http.StatusOK, getResp.StatusCode)

			var retrievedVacancy *vacancies.Vacancy
			err = json.NewDecoder(getResp.Body).Decode(&retrievedVacancy)
			assert.NoError(t, err)
			assert.Equal(t, updateDto.Salary, retrievedVacancy.Salary)
			assert.Equal(t, updateDto.Title, retrievedVacancy.Title)
		})

		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			updateDto := vacancies.UpdateVacancyDto{
				Title:  "President",
				Salary: 2000000,
			}
			jsonUpdateDto, _ := json.Marshal(updateDto)

			updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/ABABABA", server.URL), bytes.NewBuffer(jsonUpdateDto))
			updateResp, err := http.DefaultClient.Do(updateReq)
			assert.NoError(t, err)

			defer updateResp.Body.Close()
			assert.Equal(t, http.StatusNotFound, updateResp.StatusCode)

		})

		t.Run("invalid body", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/ABABABA", server.URL), bytes.NewBuffer([]byte("bebra")))
			updateResp, err := http.DefaultClient.Do(updateReq)
			assert.NoError(t, err)

			defer updateResp.Body.Close()
			assert.Equal(t, http.StatusBadRequest, updateResp.StatusCode)
		})
	})
	t.Run("Delete", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			createDto := vacancies.CreateVacancyDto{
				Title:  "Software Engineer",
				Salary: 100000,
			}

			jsonData, _ := json.Marshal(createDto)
			req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(jsonData))
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			var respBody struct{ Id string }
			err = json.NewDecoder(resp.Body).Decode(&respBody)
			assert.NoError(t, err)

			deleteReq, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
			deleteResp, err := http.DefaultClient.Do(deleteReq)
			assert.NoError(t, err)

			defer deleteResp.Body.Close()
			assert.Equal(t, http.StatusNoContent, deleteResp.StatusCode)

			getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
			getResp, err := http.DefaultClient.Do(getReq)
			defer getResp.Body.Close()
			assert.NoError(t, err)
			assert.Equal(t, http.StatusNotFound, getResp.StatusCode)
		})

		t.Run("not found", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)
			memcached := apps.BuildMemcached()

			repo := pg.NewVacancyPGRepo(Database.Db)
			cacheRepo := memcached2.NewVacancyMemcachedRepo(memcached)
			svc := vacancies.NewVacancySvc(repo)
			cacheSvc := vacancies.NewCacheVacancySvc(cacheRepo, svc)
			controller := vacancies_controller.NewVacancyController(cacheSvc)
			router := vacancies_controller.GetRouter(controller)
			server := httptest.NewServer(router)
			defer server.Close()

			deleteReq, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/OAAOAOAOA", server.URL), nil)
			deleteResp, err := http.DefaultClient.Do(deleteReq)
			defer deleteResp.Body.Close()
			assert.NoError(t, err)

			assert.Equal(t, http.StatusNotFound, deleteResp.StatusCode)
		})
	})
}
