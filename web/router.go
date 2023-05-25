package web

import (
	"crud-server/usersstorage"
	"net/http"

	"github.com/gorilla/mux"
)

// NewPeopleStoreRouter router
func NewPeopleStoreRouter(dh *DataHandler, dbu *usersstorage.DBUHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth/{provider}/callback", dh.FCallback).Methods(http.MethodGet)
	router.HandleFunc("/logout/{provider}", dh.Logout).Methods(http.MethodGet)
	router.HandleFunc("/auth/{provider}", dh.Auth).Methods(http.MethodGet)
	router.HandleFunc("/auth", dh.DefPageAuth).Methods(http.MethodGet)

	// router.HandleFunc("/users", .AddUs).Methods(http.MethodGet)

	router.HandleFunc("/", dh.MainPage).Methods(http.MethodGet)
	router.HandleFunc("/man/all", dh.GetAllMan).Methods(http.MethodGet)
	router.HandleFunc("/man/{manID}", dh.GetManByID).Methods(http.MethodGet)
	router.HandleFunc("/man", dh.CreateMan).Methods(http.MethodPost)
	router.HandleFunc("/man/{manID}", dh.UpdateMan).Methods(http.MethodPut)
	router.HandleFunc("/man/{manID}", dh.DeleteMan).Methods(http.MethodDelete)

	return router
}
