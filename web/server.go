package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer(port string, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: router,
	}
}
