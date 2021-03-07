package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewPeopleStoreRouter router
func NewPeopleStoreRouter(dh *DataHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", dh.MainPage).Methods(http.MethodGet)
	router.HandleFunc("/man/all", dh.GetAllMan).Methods(http.MethodGet)
	router.HandleFunc("/man/{manID}", dh.GetManByID).Methods(http.MethodGet)
	router.HandleFunc("/man", dh.CreateMan).Methods(http.MethodPost)
	router.HandleFunc("/man/{manID}", dh.UpdateMan).Methods(http.MethodPut)
	router.HandleFunc("/man/{manID}", dh.DeleteMan).Methods(http.MethodDelete)

	return router
}
