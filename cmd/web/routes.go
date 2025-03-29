package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, makeResponseJSON)

	//dynamicMiddleware := alice.New()

	mux := pat.New()

	// USER
	mux.Post("/api/users", http.HandlerFunc(app.userHandler.Create))
	mux.Get("/api/users", http.HandlerFunc(app.userHandler.GetAll))
	mux.Get("/api/users/:id", http.HandlerFunc(app.userHandler.GetByID))
	mux.Put("/api/users", http.HandlerFunc(app.userHandler.Update))
	mux.Del("/api/users/:id", http.HandlerFunc(app.userHandler.Delete))

	// INCIDENT
	mux.Post("/api/incidents", http.HandlerFunc(app.incidentHandler.Create))
	mux.Get("/api/incidents", http.HandlerFunc(app.incidentHandler.GetAll))
	mux.Get("/api/incidents/:id", http.HandlerFunc(app.incidentHandler.GetByID))
	mux.Put("/api/incidents", http.HandlerFunc(app.incidentHandler.Update))
	mux.Del("/api/incidents/:id", http.HandlerFunc(app.incidentHandler.Delete))

	// EDUCATION
	mux.Post("/api/education", http.HandlerFunc(app.educationHandler.Create))
	mux.Get("/api/education", http.HandlerFunc(app.educationHandler.GetAll))
	mux.Get("/api/education/:id", http.HandlerFunc(app.educationHandler.GetByID))
	mux.Put("/api/education", http.HandlerFunc(app.educationHandler.Update))
	mux.Del("/api/education/:id", http.HandlerFunc(app.educationHandler.Delete))

	// EMERGENCY
	mux.Post("/api/emergency", http.HandlerFunc(app.emergencyHandler.Create))
	mux.Get("/api/emergency", http.HandlerFunc(app.emergencyHandler.GetAll))
	mux.Get("/api/emergency/:id", http.HandlerFunc(app.emergencyHandler.GetByID))
	mux.Del("/api/emergency/:id", http.HandlerFunc(app.emergencyHandler.Delete))

	// NEWS
	mux.Post("/api/news", http.HandlerFunc(app.newsHandler.Create))
	mux.Get("/api/news", http.HandlerFunc(app.newsHandler.GetAll))
	mux.Get("/api/news/:id", http.HandlerFunc(app.newsHandler.GetByID))
	mux.Put("/api/news", http.HandlerFunc(app.newsHandler.Update))
	mux.Del("/api/news/:id", http.HandlerFunc(app.newsHandler.Delete))

	// MESSAGES
	mux.Post("/api/messages", http.HandlerFunc(app.messageHandler.Create))
	mux.Get("/api/messages", http.HandlerFunc(app.messageHandler.GetAll))
	mux.Get("/api/messages/:id", http.HandlerFunc(app.messageHandler.GetByID))
	mux.Del("/api/messages/:id", http.HandlerFunc(app.messageHandler.Delete))

	// NOTIFY TOKEN
	mux.Post("/api/notify-tokens", http.HandlerFunc(app.notifyTokenHandler.Create))
	mux.Get("/api/notify-tokens", http.HandlerFunc(app.notifyTokenHandler.GetAll))
	mux.Del("/api/notify-tokens/:id", http.HandlerFunc(app.notifyTokenHandler.Delete))

	// NOTIFY HISTORY
	mux.Post("/api/notify-history", http.HandlerFunc(app.notifyHistoryHandler.Create))
	mux.Get("/api/notify-history", http.HandlerFunc(app.notifyHistoryHandler.GetAll))
	mux.Get("/api/notify-history/:id", http.HandlerFunc(app.notifyHistoryHandler.GetByID))

	// POLICE
	mux.Post("/api/police-department", http.HandlerFunc(app.policeHandler.Create))
	mux.Get("/api/police-department", http.HandlerFunc(app.policeHandler.GetAll))
	mux.Get("/api/police-department/:id", http.HandlerFunc(app.policeHandler.GetByID))
	mux.Put("/api/police-department", http.HandlerFunc(app.policeHandler.Update))
	mux.Del("/api/police-department/:id", http.HandlerFunc(app.policeHandler.Delete))

	// NOTIFY
	mux.Post("/notify", http.HandlerFunc(app.fcmHandler.NotifyChange))
	mux.Post("/notify/token/create", http.HandlerFunc(app.fcmHandler.CreateToken))
	mux.Del("/notify/token/:id", http.HandlerFunc(app.fcmHandler.DeleteToken))
	mux.Post("/notify/history", http.HandlerFunc(app.fcmHandler.ShowNotifyHistory))
	mux.Del("/notify/history/:id", http.HandlerFunc(app.fcmHandler.DeleteNotifyHistory))

	// WEBSOCKET
	mux.Get("/ws", http.HandlerFunc(app.WebSocketHandler))

	return standardMiddleware.Then(mux)
}
