package candidate_vacancies_controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter(controller CandidateVacanciesHttpHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Create).Methods(http.MethodPost)
	r.HandleFunc("/", controller.DeleteResponseForVacancy).Methods(http.MethodDelete)
	r.HandleFunc("/candidates", controller.GetCandidatesByVacancyId).Methods(http.MethodGet)
	r.HandleFunc("/vacancies", controller.GetVacanciesByCandidate).Methods(http.MethodGet)

	return r

}
