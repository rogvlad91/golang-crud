package candidate_controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"homework-7/internal/svc/candidates"
	"homework-7/internal/svc/candidates/repo/pg"
	"net/http"
)

type CandidateController struct {
	svc candidates.CandidateProcessor
}

func NewCandidateController(svc candidates.CandidateProcessor) *CandidateController {
	return &CandidateController{svc: svc}
}

func (c *CandidateController) Create(w http.ResponseWriter, r *http.Request) {
	var dto candidates.CreateCandidateDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
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

func (c *CandidateController) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	candidate, err := c.svc.GetById(r.Context(), id)
	if err != nil {
		if err.Error() == "candidate not found" {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candidate)
}

func (c *CandidateController) GetAll(w http.ResponseWriter, r *http.Request) {
	foundCandidates, err := c.svc.GetAll(r.Context())
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundCandidates)
}

func (c *CandidateController) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var dto candidates.UpdateCandidateDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = c.svc.Update(r.Context(), id, dto)

	if err != nil {
		if errors.Is(err, pg.CandidateNotFoundError) {
			http.Error(w, "candidate not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CandidateController) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := c.svc.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pg.CandidateNotFoundError) {
			http.Error(w, "candidate not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type CandidateHttpHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
