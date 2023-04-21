package vacancies_controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter(controller VacancyHTTPHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Create).Methods(http.MethodPost)
	r.HandleFunc("/{id}", controller.GetById).Methods(http.MethodGet)
	r.HandleFunc("/", controller.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id}", controller.Update).Methods(http.MethodPut)
	r.HandleFunc("/{id}", controller.Delete).Methods(http.MethodDelete)

	return r
}
