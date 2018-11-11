package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

//go:generate moq -out router_mock.go . Router
type Router interface {
	Methods(methods ...string) *mux.Route
	Handle(path string, handler http.Handler) *mux.Route
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
