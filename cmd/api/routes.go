package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.GET("/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/", app.landing)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.insertMovie)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovie)
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getOneMovie)
	return app.enableCORS(router)
}
