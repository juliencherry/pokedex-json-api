package main

import (
	"fmt"
	"log"
	"net/http"

	"pokedex.juliencherry.net/env"
	"pokedex.juliencherry.net/server"
)

func main() {
	cfg, err := env.LoadEnvironment()
	if err != nil {
		log.Fatalf("could not load environment variables: %s", err.Error())
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on %s\n", port)
	if err := http.ListenAndServe(port, server.Server{}); err != nil {
		log.Fatal(err)
	}
}
