package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, makeResponseJSON)

	//dynamicMiddleware := alice.New()

	mux := pat.New()

	// Swagger docs
	mux.Get("/swagger/", httpSwagger.WrapHandler)

	// USER
	// @Tags Users
	// @Router /api/users [post]
	mux.Post("/api/users", http.HandlerFunc(app.userHandler.Create))
	// @Router /api/users [get]
	mux.Get("/api/users", http.HandlerFunc(app.userHandler.GetAll))
	// @Router /api/users/{id} [get]
	mux.Get("/api/users/:id", http.HandlerFunc(app.userHandler.GetByID))
	// @Router /api/users [put]
	mux.Put("/api/users", http.HandlerFunc(app.userHandler.Update))
	// @Router /api/users/{id} [delete]
	mux.Del("/api/users/:id", http.HandlerFunc(app.userHandler.Delete))

	// INCIDENT
	// @Tags Incidents
	mux.Post("/api/incidents", http.HandlerFunc(app.incidentHandler.Create))
	mux.Get("/api/incidents", http.HandlerFunc(app.incidentHandler.GetAll))
	mux.Get("/api/incidents/:id", http.HandlerFunc(app.incidentHandler.GetByID))
	mux.Put("/api/incidents", http.HandlerFunc(app.incidentHandler.Update))
	mux.Del("/api/incidents/:id", http.HandlerFunc(app.incidentHandler.Delete))

	// EDUCATION
	// @Tags Education
	mux.Post("/api/education", http.HandlerFunc(app.educationHandler.Create))
	mux.Get("/api/education", http.HandlerFunc(app.educationHandler.GetAll))
	mux.Get("/api/education/:id", http.HandlerFunc(app.educationHandler.GetByID))
	mux.Put("/api/education", http.HandlerFunc(app.educationHandler.Update))
	mux.Del("/api/education/:id", http.HandlerFunc(app.educationHandler.Delete))

	// EMERGENCY
	// @Tags Emergency
	mux.Post("/api/emergency", http.HandlerFunc(app.emergencyHandler.Create))
	mux.Get("/api/emergency", http.HandlerFunc(app.emergencyHandler.GetAll))
	mux.Get("/api/emergency/:id", http.HandlerFunc(app.emergencyHandler.GetByID))
	mux.Del("/api/emergency/:id", http.HandlerFunc(app.emergencyHandler.Delete))

	// NEWS
	// @Tags News
	mux.Post("/api/news", http.HandlerFunc(app.newsHandler.Create))
	mux.Get("/api/news", http.HandlerFunc(app.newsHandler.GetAll))
	mux.Get("/api/news/:id", http.HandlerFunc(app.newsHandler.GetByID))
	mux.Put("/api/news", http.HandlerFunc(app.newsHandler.Update))
	mux.Del("/api/news/:id", http.HandlerFunc(app.newsHandler.Delete))

	// MESSAGES
	// @Tags Messages
	mux.Post("/api/messages", http.HandlerFunc(app.messageHandler.Create))
	mux.Get("/api/messages", http.HandlerFunc(app.messageHandler.GetAll))
	mux.Get("/api/messages/:id", http.HandlerFunc(app.messageHandler.GetByID))
	mux.Del("/api/messages/:id", http.HandlerFunc(app.messageHandler.Delete))

	// POLICE
	// @Tags Police
	mux.Post("/api/police-department", http.HandlerFunc(app.policeHandler.Create))
	mux.Get("/api/police-department", http.HandlerFunc(app.policeHandler.GetAll))
	mux.Get("/api/police-department/:id", http.HandlerFunc(app.policeHandler.GetByID))
	mux.Put("/api/police-department", http.HandlerFunc(app.policeHandler.Update))
	mux.Del("/api/police-department/:id", http.HandlerFunc(app.policeHandler.Delete))

	// NOTIFY (FCM)
	// @Tags Notifications
	mux.Post("/notify", http.HandlerFunc(app.fcmHandler.NotifyChange))
	mux.Post("/notify/token/create", http.HandlerFunc(app.fcmHandler.CreateToken))
	mux.Del("/notify/token/:id", http.HandlerFunc(app.fcmHandler.DeleteToken))
	mux.Post("/notify/history", http.HandlerFunc(app.fcmHandler.ShowNotifyHistory))
	mux.Del("/notify/history/:id", http.HandlerFunc(app.fcmHandler.DeleteNotifyHistory))

	// ALERTS
	// @Tags Alerts
	mux.Post("/api/alerts", http.HandlerFunc(app.alertHandler.Create))
	mux.Get("/api/alerts", http.HandlerFunc(app.alertHandler.GetAll))
	mux.Get("/api/alerts/:id", http.HandlerFunc(app.alertHandler.GetByID))
	mux.Put("/api/alerts", http.HandlerFunc(app.alertHandler.Update))
	mux.Del("/api/alerts/:id", http.HandlerFunc(app.alertHandler.Delete))

	// WEBSOCKET
	// @Tags WebSocket
	// @Description Подключение WebSocket для реального времени по пути /ws
	mux.Get("/ws", http.HandlerFunc(app.WebSocketHandler))

	// @Tags WebSocket
	// @Description Подключение WebSocket чата с AI по пути /ws/ai
	mux.Get("/ws/ai", http.HandlerFunc(app.WebSocketAIHandler))

	return standardMiddleware.Then(mux)
}
