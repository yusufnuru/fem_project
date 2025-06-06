package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/yusufnuru/femProject/internal/app"
)

func SetupRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	return r
}
