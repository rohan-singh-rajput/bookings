package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/rohan-singh-rajput/bookings/pkg/config"
	"github.com/rohan-singh-rajput/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)

	return router

}
