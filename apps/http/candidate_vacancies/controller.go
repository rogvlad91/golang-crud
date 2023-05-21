package candidate_vacancies_controller

import (
	"encoding/json"
	"errors"
	"golang-crud/internal/svc/candidate_vacancies"
	"golang-crud/internal/svc/candidate_vacancies/repo/pg"
	"net/http"
)

type CandidateVacanciesHttpHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	DeleteResponseForVacancy(w http.ResponseWriter, r *http.Request)
	GetCandidatesByVacancyId(w http.ResponseWriter, r *http.Request)
	GetVacanciesByCandidate(w http.ResponseWriter, r *http.Request)
}

type CandidateVacanciesController struct {
	svc candidate_vacancies.CandidateVacanciesProcessor
}

func NewCandidateVacanciesController(svc candidate_vacancies.CandidateVacanciesProcessor) *CandidateVacanciesController {
	return &CandidateVacanciesController{svc: svc}
}

func (c *CandidateVacanciesController) Create(w http.ResponseWriter, r *http.Request) {
	var dto candidate_vacancies.CreateDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if dto.CandidateId == "" || dto.VacancyId == "" {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := c.svc.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{ Id string }{Id: id})
}

func (c *CandidateVacanciesController) DeleteResponseForVacancy(w http.ResponseWriter, r *http.Request) {
	vacancyId := r.URL.Query().Get("vacancy_id")
	candidateId := r.URL.Query().Get("candidate_id")

	err := c.svc.DeleteResponseForVacancy(r.Context(), vacancyId, candidateId)
	if err != nil {
		if errors.Is(err, pg.ResponseNotFoundError) {
			http.Error(w, "response not found", http.StatusNotFound)
		} else {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CandidateVacanciesController) GetCandidatesByVacancyId(w http.ResponseWriter, r *http.Request) {
	vacancyId := r.URL.Query().Get("vacancy_id")

	result, err := c.svc.GetCandidatesByVacancyId(r.Context(), vacancyId)

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if len(result) == 0 {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (c *CandidateVacanciesController) GetVacanciesByCandidate(w http.ResponseWriter, r *http.Request) {
	candidateId := r.URL.Query().Get("candidate_id")

	result, err := c.svc.GetVacanciesByCandidate(r.Context(), candidateId)

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if len(result) == 0 {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(result)
}
