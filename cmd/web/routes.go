package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// recover from panics, send back 500, keep the application running
	mux.Use(middleware.Recoverer)

	// any request that takes more than 60s will timeout
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	return mux
}
