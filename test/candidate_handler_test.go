//go:build integration
// +build integration

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"homework-7/apps"
	candidatesHttp "homework-7/apps/http/candidates"
	"homework-7/internal/svc/candidates"
	candidateMemcachedRepo "homework-7/internal/svc/candidates/repo/memcached"
	"homework-7/internal/svc/candidates/repo/pg"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCandidateHandler(t *testing.T) {
	t.Parallel()
	t.Run("createCandidate", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			Database.SetUp(t)
			defer Database.TearDown(t)

			repo := pg.NewCandidatePGRepo(Database.Db)
			candidateSvc := candidates.NewCandidateSvc(repo)
			memCachedRepo := apps.BuildMemcached()
			candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
			cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

			candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

			candidateRouter := candidatesHttp.GetRouter(candidatesController)
			server := httptest.NewServer(candidateRouter)
			defer server.Close()

			dto := candidates.CreateCandidateDTO{
				FullName:     "John Doe",
				Age:          14,
				WantedJob:    "janitor",
				WantedSalary: 5000,
			}

			body, _ := json.Marshal(dto)

			req, err := http.NewRequest(http.MethodPost, server.URL+"/", bytes.NewBuffer(body))
			assert.NoError(t, err)

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

			repo := pg.NewCandidatePGRepo(Database.Db)
			candidateSvc := candidates.NewCandidateSvc(repo)
			memCachedRepo := apps.BuildMemcached()
			candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
			cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

			candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

			candidateRouter := candidatesHttp.GetRouter(candidatesController)
			server := httptest.NewServer(candidateRouter)
			defer server.Close()

			invalidBody := []byte(`{"fullName": "John Doe", "age": "not an int", "wantedJob" : "janitor", "wantedSalary": 5000}`)
			req, err := http.NewRequest("POST", server.URL+"/", bytes.NewBuffer(invalidBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			assert.NoError(t, err)

			defer res.Body.Close()

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		})
		t.Run("getCandidateById", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				createDTO := candidates.CreateCandidateDTO{
					FullName:     "John Doe",
					Age:          30,
					WantedJob:    "Software Engineer",
					WantedSalary: 100000,
				}
				candidateBytes, _ := json.Marshal(createDTO)
				req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(candidateBytes))
				resp, err := http.DefaultClient.Do(req)
				assert.NoError(t, err)

				defer resp.Body.Close()
				var respBody struct{ Id string }
				err = json.NewDecoder(resp.Body).Decode(&respBody)
				assert.NoError(t, err)

				getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
				getResp, err := http.DefaultClient.Do(getReq)
				defer getResp.Body.Close()
				assert.NoError(t, err)

				assert.NotNil(t, getResp.Body)

				assert.Equal(t, http.StatusOK, getResp.StatusCode)
				var retrievedCandidate *candidates.Candidate
				err = json.NewDecoder(getResp.Body).Decode(&retrievedCandidate)
				assert.NoError(t, err)

				assert.Equal(t, createDTO.FullName, retrievedCandidate.FullName)
				assert.Equal(t, createDTO.Age, retrievedCandidate.Age)
				assert.Equal(t, createDTO.WantedJob, retrievedCandidate.WantedJob)
				assert.Equal(t, createDTO.WantedSalary, retrievedCandidate.WantedSalary)
			})
			t.Run("not found", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/hop_hey", server.URL), nil)
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

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				for i := 0; i < 3; i++ {
					createDTO := candidates.CreateCandidateDTO{
						FullName:     fmt.Sprintf("John Doe %d", i),
						Age:          30 + i,
						WantedJob:    "Software Engineer",
						WantedSalary: 100000,
					}
					candidateBytes, _ := json.Marshal(createDTO)
					req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(candidateBytes))
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
				var candidates []*candidates.Candidate
				err = json.NewDecoder(resp.Body).Decode(&candidates)
				assert.NoError(t, err)

				assert.Equal(t, 3, len(candidates))

				for i := 0; i < 3; i++ {
					assert.Equal(t, fmt.Sprintf("John Doe %d", i), candidates[i].FullName)
					assert.Equal(t, 30+i, candidates[i].Age)
					assert.Equal(t, "Software Engineer", candidates[i].WantedJob)
					assert.Equal(t, 100000, candidates[i].WantedSalary)
				}
			})
		})

		t.Run("Update", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				createDTO := candidates.CreateCandidateDTO{
					FullName:     "John Doe",
					Age:          30,
					WantedJob:    "Software Engineer",
					WantedSalary: 100000,
				}
				candidateBytes, _ := json.Marshal(createDTO)
				req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(candidateBytes))
				resp, err := http.DefaultClient.Do(req)
				assert.NoError(t, err)

				defer resp.Body.Close()
				var respBody struct{ Id string }
				err = json.NewDecoder(resp.Body).Decode(&respBody)
				assert.NoError(t, err)

				updateDTO := candidates.UpdateCandidateDto{
					WantedJob:    "Data Scientist",
					WantedSalary: 150000,
				}
				updateBytes, _ := json.Marshal(updateDTO)
				updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/%s", server.URL, respBody.Id), bytes.NewReader(updateBytes))
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
				var retrievedCandidate *candidates.Candidate
				err = json.NewDecoder(getResp.Body).Decode(&retrievedCandidate)
				assert.NoError(t, err)

				assert.Equal(t, updateDTO.WantedJob, retrievedCandidate.WantedJob)
				assert.Equal(t, updateDTO.WantedSalary, retrievedCandidate.WantedSalary)
			})
			t.Run("invalid body", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/1", server.URL), bytes.NewReader([]byte("dddd")))
				updateResp, _ := http.DefaultClient.Do(updateReq)
				defer updateResp.Body.Close()
				assert.Equal(t, http.StatusBadRequest, updateResp.StatusCode)

			})

			t.Run("not found", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				updateDTO := candidates.UpdateCandidateDto{
					WantedJob:    "Data Scientist",
					WantedSalary: 150000,
				}
				updateBytes, _ := json.Marshal(updateDTO)
				updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("%s/non-existing-id", server.URL), bytes.NewReader(updateBytes))
				updateResp, err := http.DefaultClient.Do(updateReq)
				assert.NoError(t, err)

				defer updateResp.Body.Close()

				assert.Equal(t, http.StatusNotFound, updateResp.StatusCode)
			})
		})

		t.Run("Delete", func(t *testing.T) {
			t.Run("success", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				createDTO := candidates.CreateCandidateDTO{
					FullName:     "John Doe",
					Age:          30,
					WantedJob:    "Software Engineer",
					WantedSalary: 100000,
				}
				candidateBytes, _ := json.Marshal(createDTO)
				req, _ := http.NewRequest("POST", server.URL+"/", bytes.NewReader(candidateBytes))
				resp, err := http.DefaultClient.Do(req)
				assert.NoError(t, err)

				defer resp.Body.Close()
				var respBody struct{ Id string }
				err = json.NewDecoder(resp.Body).Decode(&respBody)
				assert.NoError(t, err)

				deleteReq, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
				deleteResp, err := http.DefaultClient.Do(deleteReq)
				defer deleteResp.Body.Close()
				assert.NoError(t, err)

				assert.Equal(t, http.StatusNoContent, deleteResp.StatusCode)

				getReq, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", server.URL, respBody.Id), nil)
				getResp, err := http.DefaultClient.Do(getReq)
				assert.NoError(t, err)

				assert.Equal(t, http.StatusNotFound, getResp.StatusCode)
			})
			t.Run("not found", func(t *testing.T) {
				Database.SetUp(t)
				defer Database.TearDown(t)

				repo := pg.NewCandidatePGRepo(Database.Db)
				candidateSvc := candidates.NewCandidateSvc(repo)
				memCachedRepo := apps.BuildMemcached()
				candidateCacheRepo := candidateMemcachedRepo.NewCandidatesMemcachedRepo(memCachedRepo)
				cacheCandidateSvc := candidates.NewCacheCandidateSvc(candidateCacheRepo, candidateSvc)

				candidatesController := candidatesHttp.NewCandidateController(cacheCandidateSvc)

				candidateRouter := candidatesHttp.GetRouter(candidatesController)
				server := httptest.NewServer(candidateRouter)
				defer server.Close()

				deleteReq, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", server.URL, "non-existing-id"), nil)
				deleteResp, err := http.DefaultClient.Do(deleteReq)
				defer deleteResp.Body.Close()
				assert.NoError(t, err)

				assert.Equal(t, http.StatusNotFound, deleteResp.StatusCode)

			})
		})
	})
}
