package handlers_test

import (
	"groupie-tracker/nyeltay/algaliyev/internal/handlers"
	"groupie-tracker/nyeltay/algaliyev/internal/requestapi"
	"groupie-tracker/nyeltay/algaliyev/internal/templates"
	"log"
	"testing"
)

func newTestApplication(t *testing.T) *handlers.Handler {
	templateCache, err := templates.NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	apiRequest := requestapi.New()

	handler := handlers.NewApplication(templateCache, apiRequest)
	return handler
}
