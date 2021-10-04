# Pokémon JSON API

## Task

Using https://pokeapi.co/ and https://jsonapi.org/

1. User requests `GET api/pokemon/:pokemon_name` from API
2. API responds with pokemon data as per jsonapi.org
3. Nice to Have: User requests `GET api/pokemon`, and it also responds with pokemon data as per jsonapi.org (with a collection)

## Prerequisites

* [Golang](https://golang.org/doc/install)

## Environment variables

These can be configured as regular environment variables or a root-level `.env` file.

* `PORT`, defaults to 8081

## Getting started

1. `go get -u`
2. `go run main.go`
