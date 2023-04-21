package candidate_controller

import (
	"github.com/gorilla/mux"
)

func GetRouter(controller CandidateHttpHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Create).Methods("POST")
	r.HandleFunc("/", controller.GetAll).Methods("GET")
	r.HandleFunc("/{id}", controller.GetById).Methods("GET")
	r.HandleFunc("/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/{id}", controller.Delete).Methods("DELETE")

	return r
}
