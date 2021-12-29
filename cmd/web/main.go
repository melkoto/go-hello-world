package main

import (
	"fmt"
	"go-hello-world/pkg/config"
	"go-hello-world/pkg/handlers"
	"go-hello-world/pkg/render"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	app.TemplateCache = tc
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting port on %s", port))

	_ = http.ListenAndServe(port, nil)
}
