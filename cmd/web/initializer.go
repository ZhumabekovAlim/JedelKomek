package main

import (
	"JedelKomek/internal/handlers"
	"JedelKomek/internal/repositories"
	"JedelKomek/internal/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type application struct {
	errorLog             *log.Logger
	infoLog              *log.Logger
	wsManager            *WebSocketManager
	userHandler          *handlers.UserHandler
	incidentHandler      *handlers.IncidentHandler
	educationHandler     *handlers.EducationHandler
	emergencyHandler     *handlers.EmergencyHandler
	newsHandler          *handlers.NewsHandler
	messageHandler       *handlers.MessageHandler
	notifyTokenHandler   *handlers.NotifyTokenHandler
	notifyHistoryHandler *handlers.NotifyHistoryHandler
	policeHandler        *handlers.PoliceDepartmentHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {
	userRepo := &repositories.UserRepository{Db: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	incidentRepo := &repositories.IncidentRepository{Db: db}
	incidentService := &services.IncidentService{Repo: incidentRepo}
	incidentHandler := &handlers.IncidentHandler{Service: incidentService}

	educationRepo := &repositories.EducationRepository{Db: db}
	educationService := &services.EducationService{Repo: educationRepo}
	educationHandler := &handlers.EducationHandler{Service: educationService}

	emergencyRepo := &repositories.EmergencyRepository{Db: db}
	emergencyService := &services.EmergencyService{Repo: emergencyRepo}
	emergencyHandler := &handlers.EmergencyHandler{Service: emergencyService}

	newsRepo := &repositories.NewsRepository{Db: db}
	newsService := &services.NewsService{Repo: newsRepo}
	newsHandler := &handlers.NewsHandler{Service: newsService}

	messageRepo := &repositories.MessageRepository{Db: db}
	messageService := &services.MessageService{Repo: messageRepo}
	messageHandler := &handlers.MessageHandler{Service: messageService}

	notifyTokenRepo := &repositories.NotifyTokenRepository{Db: db}
	notifyTokenService := &services.NotifyTokenService{Repo: notifyTokenRepo}
	notifyTokenHandler := &handlers.NotifyTokenHandler{Service: notifyTokenService}

	notifyHistoryRepo := &repositories.NotifyHistoryRepository{Db: db}
	notifyHistoryService := &services.NotifyHistoryService{Repo: notifyHistoryRepo}
	notifyHistoryHandler := &handlers.NotifyHistoryHandler{Service: notifyHistoryService}

	policeRepo := &repositories.PoliceDepartmentRepository{Db: db}
	policeService := &services.PoliceDepartmentService{Repo: policeRepo}
	policeHandler := &handlers.PoliceDepartmentHandler{Service: policeService}

	return &application{
		errorLog:             errorLog,
		infoLog:              infoLog,
		wsManager:            NewWebSocketManager(),
		userHandler:          userHandler,
		incidentHandler:      incidentHandler,
		educationHandler:     educationHandler,
		emergencyHandler:     emergencyHandler,
		newsHandler:          newsHandler,
		messageHandler:       messageHandler,
		notifyTokenHandler:   notifyTokenHandler,
		notifyHistoryHandler: notifyHistoryHandler,
		policeHandler:        policeHandler,
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("%v", err)
		panic("failed to connect to database")
		return nil, err
	}
	db.SetMaxIdleConns(35)
	if err = db.Ping(); err != nil {
		log.Printf("%v", err)
		panic("failed to ping the database")
		return nil, err
	}
	fmt.Println("successfully connected")

	return db, nil
}

func addSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		next.ServeHTTP(w, r)
	})
}
