package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rohan-singh-rajput/bookings/pkg/config"
	"github.com/rohan-singh-rajput/bookings/pkg/handlers"
	"github.com/rohan-singh-rajput/bookings/pkg/render"
)

const PortNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("listening at port", PortNumber)

	_ = http.ListenAndServe(PortNumber, nil)
}
