package server

import (
	"net/http"
)

type Server struct{}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/api/pokemon", s.handlePokemonRoute())

	router.ServeHTTP(w, r)
}

func (s Server) handlePokemonRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}
