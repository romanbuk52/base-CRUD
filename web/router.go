package web

import (
	"github.com/gorilla/mux"
)

// NewPeopleStoreRouter router
func NewPeopleStoreRouter(handler *DataHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.MainPage).Methods("GET")
	router.HandleFunc("/man", handler.GetMan).Methods("GET")
	router.HandleFunc("/man/{manID}", handler.GetManById).Methods("GET")
	//router.HandleFunc("/man/", CreateMan).Methods("POST")
	//router.HandleFunc("/man/{manID}", web.UpdateMan).Methods("PUT")
	//router.HandleFunc("/man/{manID}", web.DeleteMan).Methods("DELETE")

	return router
}

// routes.PeopleStoreRoutes(r)
// http.Handle("/", r)
