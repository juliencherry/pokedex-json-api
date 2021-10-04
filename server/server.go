package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pokedex.juliencherry.net/pokeapi"
)

type Server struct{}

type PokemonResponse struct {
	Data   pokeapi.Pokemon  `json:"data,omitempty"`
	Errors []ResponseErrors `json:"errors,omitempty"`
}

type ResponseErrors struct {
	Status string `json:"status"`
	Title  string `json:"title"`
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/api/pokemon", s.handlePokemonRoute())

	router.ServeHTTP(w, r)
}

func (s Server) handlePokemonRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := preparePokemonResponse(r)

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(res)
	}
}

func preparePokemonResponse(r *http.Request) *PokemonResponse {
	response := &PokemonResponse{}

	if r.Method != http.MethodGet {
		response.Errors = append(response.Errors, ResponseErrors{
			Status: fmt.Sprint(http.StatusMethodNotAllowed),
			Title:  http.StatusText(http.StatusMethodNotAllowed),
		})
		return response
	}

	if r.Header.Get("Content-Type") != "application/vnd.api+json" {
		response.Errors = append(response.Errors, ResponseErrors{
			Status: fmt.Sprint(http.StatusUnsupportedMediaType),
			Title:  http.StatusText(http.StatusUnsupportedMediaType),
		})
		return response
	}

	client := &pokeapi.Client{}
	pokemon, err := client.Ditto()
	if err != nil {
		response.Errors = append(response.Errors, ResponseErrors{
			Status: fmt.Sprint(http.StatusNotFound),
			Title:  err.Error(),
		})
	}

	response.Data = pokemon
	return response
}
