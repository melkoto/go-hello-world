package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go-hello-world/pkg/config"
	"go-hello-world/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
