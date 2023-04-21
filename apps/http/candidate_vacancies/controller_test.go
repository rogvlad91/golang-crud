package candidate_vacancies_controller

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"homework-7/internal/svc/candidate_vacancies/repo/pg"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateCandidateVacancy(t *testing.T) {
	mockCreateDto := []byte(`{"candidate_id": "1", "vacancy_id": "2"}`)
	mockId := "123"
	t.Run("CreateResponse", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/candidate-vacancies", bytes.NewBuffer(mockCreateDto))
			s.mockSvc.EXPECT().Create(r.Context(), gomock.Any()).Return(mockId, nil)

			s.controller.Create(w, r)

			assert.Equal(t, http.StatusCreated, w.Code)
			assert.Equal(t, "{\"Id\":\"123\"}\n", w.Body.String())
		})

		t.Run("Invalid request payload", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/candidate-vacancies", strings.NewReader("invalid payload"))

			s.controller.Create(w, r)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "invalid request payload\n", w.Body.String())
		})

		t.Run("Internal server error", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/candidate-vacancies", bytes.NewBuffer(mockCreateDto))
			s.mockSvc.EXPECT().Create(r.Context(), gomock.Any()).Return("", assert.AnError)

			s.controller.Create(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.Equal(t, "server error\n", w.Body.String())
		})
	})

	t.Run("DeleteResponse", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockVacancyID := "123"
			mockCandidateID := "456"
			s := setUp(t)
			defer s.tearDown()

			r, _ := http.NewRequest(http.MethodDelete, "/candidate-vacancies?vacancy_id="+mockVacancyID+"&candidate_id="+mockCandidateID, nil)
			w := httptest.NewRecorder()

			s.mockSvc.EXPECT().DeleteResponseForVacancy(r.Context(), mockVacancyID, mockCandidateID).Return(nil)

			s.controller.DeleteResponseForVacancy(w, r)

			assert.Equal(t, http.StatusNoContent, w.Code)
		})

		t.Run("Response not found", func(t *testing.T) {
			mockVacancyID := "123"
			mockCandidateID := "456"
			s := setUp(t)
			defer s.tearDown()

			r, _ := http.NewRequest(http.MethodDelete, "/candidate-vacancies?vacancy_id="+mockVacancyID+"&candidate_id="+mockCandidateID, nil)
			w := httptest.NewRecorder()

			s.mockSvc.EXPECT().DeleteResponseForVacancy(r.Context(), mockVacancyID, mockCandidateID).Return(pg.ResponseNotFoundError)

			s.controller.DeleteResponseForVacancy(w, r)

			assert.Equal(t, http.StatusNotFound, w.Code)
		})

		t.Run("Internal server error", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()
			mockVacancyID := "123"
			mockCandidateID := "456"

			r, _ := http.NewRequest(http.MethodDelete, "/candidate-vacancies?vacancy_id="+mockVacancyID+"&candidate_id="+mockCandidateID, nil)
			w := httptest.NewRecorder()

			s.mockSvc.EXPECT().DeleteResponseForVacancy(r.Context(), mockVacancyID, mockCandidateID).Return(assert.AnError)

			s.controller.DeleteResponseForVacancy(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
	t.Run("GetCandidatesForVacancy", func(t *testing.T) {
		mockVacancyID := "123"
		mockCandidate1 := `{"ID":"1","Name":"Alice"}`
		mockCandidate2 := `{"ID":"2","Name":"Bob"}`

		mockResult := []*string{&mockCandidate1, &mockCandidate2}

		t.Run("Success", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/candidate-vacancies/candidates?vacancy_id="+mockVacancyID, nil)
			s.mockSvc.EXPECT().GetCandidatesByVacancyId(r.Context(), mockVacancyID).Return(mockResult, nil)

			s.controller.GetCandidatesByVacancyId(w, r)

			assert.Equal(t, http.StatusOK, w.Code)

			var res []*string
			json.Unmarshal(w.Body.Bytes(), &res)

			assert.Equal(t, mockResult, res)
		})

		t.Run("Vacancy not found", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/candidate-vacancies/candidates?vacancy_id="+mockVacancyID, nil)
			s.mockSvc.EXPECT().GetCandidatesByVacancyId(r.Context(), mockVacancyID).Return(nil, nil)

			s.controller.GetCandidatesByVacancyId(w, r)

			assert.Equal(t, http.StatusNotFound, w.Code)
		})

		t.Run("Internal server error", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/candidate-vacancies/candidates?vacancy_id="+mockVacancyID, nil)
			s.mockSvc.EXPECT().GetCandidatesByVacancyId(r.Context(), mockVacancyID).Return(nil, assert.AnError)

			s.controller.GetCandidatesByVacancyId(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

	})

	t.Run("GetVacanciesForCandidate", func(t *testing.T) {
		mockCandidateId := "123"
		mockCandidate1 := `{"Title":"Coder","Salary":5000}`
		mockCandidate2 := `{"Title":"President","Salary":"1337"}`

		mockResult := []*string{&mockCandidate1, &mockCandidate2}

		t.Run("Success", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/candidate-vacancies/vacancies?candidate_id="+mockCandidateId, nil)
			s.mockSvc.EXPECT().GetVacanciesByCandidate(r.Context(), mockCandidateId).Return(mockResult, nil)

			s.controller.GetVacanciesByCandidate(w, r)

			assert.Equal(t, http.StatusOK, w.Code)

			var res []*string
			json.Unmarshal(w.Body.Bytes(), &res)

			assert.Equal(t, mockResult, res)
		})

		t.Run("Candidate not found", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/candidate-vacancies/vacancies?candidate_id="+mockCandidateId, nil)
			s.mockSvc.EXPECT().GetVacanciesByCandidate(r.Context(), mockCandidateId).Return(nil, nil)

			s.controller.GetVacanciesByCandidate(w, r)

			assert.Equal(t, http.StatusNotFound, w.Code)
		})

		t.Run("Internal server error", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/candidate-vacancies/vacancies?candidate_id="+mockCandidateId, nil)
			s.mockSvc.EXPECT().GetVacanciesByCandidate(r.Context(), mockCandidateId).Return(nil, assert.AnError)

			s.controller.GetVacanciesByCandidate(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

	})
}
