package vacancies_controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"homework-7/internal/svc/vacancies"
	"homework-7/internal/svc/vacancies/repo/pg"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestVacancyController(t *testing.T) {
	reqBody := []byte(`{"title": "coder", "salary": 30000}`)
	mockVacancy := &vacancies.Vacancy{
		Id:        "1488",
		Title:     "Coder",
		Salary:    30000,
		CreatedAt: "",
		UpdatedAt: "",
	}
	t.Run("Create", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {

			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/vacancies", bytes.NewBuffer(reqBody))
			s.mockSvc.EXPECT().Create(r.Context(), gomock.Any()).Return("1488", nil)
			s.controller.Create(w, r)

			assert.Equal(t, http.StatusCreated, w.Code)
			assert.Equal(t, "{\"Id\":\"1488\"}\n", w.Body.String())
		})
		t.Run("InvalidRequestPayload", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/vacancies", strings.NewReader("invalid payload"))

			s.controller.Create(w, r)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "invalid request payload\n", w.Body.String())
		})

		t.Run("ServerError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/vacancies", bytes.NewBuffer(reqBody))
			s.mockSvc.EXPECT().Create(r.Context(), gomock.Any()).Return("", assert.AnError)

			s.controller.Create(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.Equal(t, "server error\n", w.Body.String())
		})
	})
	t.Run("GetById", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/vacancies/1", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().GetById(r.Context(), "1").Return(mockVacancy, nil)

			s.controller.GetById(w, r)
			expectedBody := `
		{"Id":"1488","Title":"Coder","Salary":30000, "CreatedAt": "", "UpdatedAt": ""}`
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, expectedBody, w.Body.String())
		})

		t.Run("NotFound", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/vacancies/1", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().GetById(r.Context(), "1").Return(nil, errors.New("vacancy not found"))

			s.controller.GetById(w, r)

			assert.Equal(t, http.StatusNotFound, w.Code)
		})

		t.Run("ServerError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/vacancies/1", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().GetById(r.Context(), "1").Return(nil, assert.AnError)

			s.controller.GetById(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.Equal(t, "server error\n", w.Body.String())
		})
	})
	t.Run("GetAll", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			vacanciesRes := []*vacancies.Vacancy{
				{
					Id:     "1",
					Title:  "coder",
					Salary: 30000,
				},
				{
					Id:     "2",
					Title:  "developer",
					Salary: 50000,
				},
			}

			s.mockSvc.EXPECT().GetAll(gomock.Any()).Return(vacanciesRes, nil)

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/vacancies", nil)
			s.controller.GetAll(w, r)

			assert.Equal(t, http.StatusOK, w.Code)

			expectedBody := `[
				{"Id":"1","Title":"coder","Salary":30000, "CreatedAt": "", "UpdatedAt": ""},
				{"Id":"2", "Title":"developer","Salary":50000, "CreatedAt": "", "UpdatedAt": ""}
			]`
			assert.JSONEq(t, expectedBody, w.Body.String())
		})
		t.Run("InternalError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			s.mockSvc.EXPECT().GetAll(gomock.Any()).Return(nil, assert.AnError)

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/vacancies", nil)
			s.controller.GetAll(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.Equal(t, "server error\n", w.Body.String())
		})
	})
	t.Run("Update", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockUpdateDto := []byte(`{"Title": "coder", "Salary": 30000}`)

			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPut, "/vacancies/1", bytes.NewBuffer(mockUpdateDto))
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().Update(r.Context(), "1", gomock.Any()).Return(nil)

			s.controller.Update(w, r)

			assert.Equal(t, http.StatusNoContent, w.Code)
		})

		t.Run("InvalidPayload", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPut, "/vacancies/1488", bytes.NewBuffer([]byte(`invalid`)))
			s.controller.Update(w, r)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("ServerError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPut, "/vacancies/1488", bytes.NewBuffer(reqBody))
			r = mux.SetURLVars(r, map[string]string{"id": "1488"})
			s.mockSvc.EXPECT().Update(r.Context(), "1488", gomock.Any()).Return(assert.AnError)

			s.controller.Update(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("NotFoundError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPut, "/vacancies/1488", bytes.NewBuffer(reqBody))
			r = mux.SetURLVars(r, map[string]string{"id": "1488"})
			s.mockSvc.EXPECT().Update(r.Context(), "1488", gomock.Any()).Return(pg.VacancyNotFoundError)

			s.controller.Update(w, r)

			assert.Equal(t, http.StatusNotFound, w.Code)
			assert.Equal(t, "vacancy not found\n", w.Body.String())
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/vacancies/1", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().Delete(r.Context(), "1").Return(nil)

			s.controller.Delete(w, r)

			assert.Equal(t, http.StatusNoContent, w.Code)
		})

		t.Run("ServerError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/vacancies/1", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().Delete(r.Context(), "1").Return(assert.AnError)

			s.controller.Delete(w, r)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("NotFoundError", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/vacancies/1", nil)
			r = mux.SetURLVars(r, map[string]string{"id": "1"})
			s.mockSvc.EXPECT().Delete(r.Context(), "1").Return(pg.VacancyNotFoundError)

			s.controller.Delete(w, r)

			assert.Equal(t, "vacancy not found\n", w.Body.String())
		})
	})

}
