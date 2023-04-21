package vacancies_controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"homework-7/internal/svc/vacancies"
	"homework-7/internal/svc/vacancies/repo/pg"
	"net/http"
)

type VacancyHTTPHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type VacancyController struct {
	svc vacancies.VacancyProcessor
}

func NewVacancyController(svc vacancies.VacancyProcessor) *VacancyController {
	return &VacancyController{svc: svc}
}

func (vc *VacancyController) Create(w http.ResponseWriter, r *http.Request) {
	var createDto vacancies.CreateVacancyDto
	err := json.NewDecoder(r.Body).Decode(&createDto)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := vc.svc.Create(r.Context(), createDto)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{ Id string }{Id: id})
}

func (vc *VacancyController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vacancy, err := vc.svc.GetById(r.Context(), id)

	if err != nil {
		if err.Error() == "vacancy not found" {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vacancy)
}

func (vc *VacancyController) GetAll(w http.ResponseWriter, r *http.Request) {
	foundVacancies, err := vc.svc.GetAll(r.Context())
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(foundVacancies)
}

func (vc *VacancyController) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updateDto vacancies.UpdateVacancyDto
	err := json.NewDecoder(r.Body).Decode(&updateDto)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = vc.svc.Update(r.Context(), id, updateDto)
	if err != nil {
		if errors.Is(err, pg.VacancyNotFoundError) {
			http.Error(w, "vacancy not found", http.StatusNotFound)
		} else {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (vc *VacancyController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := vc.svc.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pg.VacancyNotFoundError) {
			http.Error(w, "vacancy not found", http.StatusNotFound)
		} else {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
