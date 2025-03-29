package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, makeResponseJSON)

	dynamicMiddleware := alice.New()

	mux := pat.New()

	// Users
	mux.Post("/api/users", http.HandlerFunc(app.userHandler.Create))
	mux.Get("/api/users", http.HandlerFunc(app.userHandler.GetAll))
	mux.Get("/api/users/:id", http.HandlerFunc(app.userHandler.GetByID))
	mux.Put("/api/users", http.HandlerFunc(app.userHandler.Update))
	mux.Del("/api/users/:id", http.HandlerFunc(app.userHandler.Delete))

	mux.Post("/api/incidents", http.HandlerFunc(app.incidentHandler.Create))
	mux.Get("/api/incidents", http.HandlerFunc(app.incidentHandler.GetAll))
	mux.Get("/api/incidents/:id", http.HandlerFunc(app.incidentHandler.GetByID))
	mux.Put("/api/incidents", http.HandlerFunc(app.incidentHandler.Update))
	mux.Del("/api/incidents/:id", http.HandlerFunc(app.incidentHandler.Delete))

	return standardMiddleware.Then(mux)
}
