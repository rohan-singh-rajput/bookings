package handlers

import (
	"net/http"

	"github.com/rohan-singh-rajput/bookings/pkg/config"
	"github.com/rohan-singh-rajput/bookings/pkg/render"
)

var Repo *Repository

// repository pattern
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
