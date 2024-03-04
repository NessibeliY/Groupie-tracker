package main

import (
	"groupie-tracker/nyeltay/algaliyev/internal/handlers"
	"groupie-tracker/nyeltay/algaliyev/internal/requestapi"
	"groupie-tracker/nyeltay/algaliyev/internal/server"
	"groupie-tracker/nyeltay/algaliyev/internal/templates"
	"log"
)

func main() {
	templateCache, err := templates.NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	apiRequest := requestapi.New()

	handler := handlers.NewApplication(templateCache, apiRequest)

	log.Fatal(server.RunServer(handler.Routes()))
}
