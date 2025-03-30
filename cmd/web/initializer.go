package main

import (
	"JedelKomek/internal/handlers"
	"JedelKomek/internal/repositories"
	"JedelKomek/internal/services"
	"context"
	"database/sql"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

type application struct {
	db               *sql.DB
	errorLog         *log.Logger
	infoLog          *log.Logger
	wsManager        *WebSocketManager
	userHandler      *handlers.UserHandler
	incidentHandler  *handlers.IncidentHandler
	educationHandler *handlers.EducationHandler
	emergencyHandler *handlers.EmergencyHandler
	fcmHandler       *handlers.FCMHandler
	newsHandler      *handlers.NewsHandler
	messageHandler   *handlers.MessageHandler
	policeHandler    *handlers.PoliceDepartmentHandler
	alertHandler     *handlers.AlertHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	ctx := context.Background()
	sa := option.WithCredentialsFile("C:\\Users\\alimz\\GolandProjects\\JedelKomek\\cmd\\web\\serviceAccountKey.json")

	firebaseApp, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "jedel-komek"}, sa)
	if err != nil {
		errorLog.Fatalf("Ошибка в нахождении приложения: %v\n", err)
	}

	fcmClient, err := firebaseApp.Messaging(ctx)
	if err != nil {
		errorLog.Fatalf("Ошибка при неверном ID устройства: %v\n", err)
	}

	fcmHandler := handlers.NewFCMHandler(fcmClient, db)

	userRepo := &repositories.UserRepository{Db: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	incidentRepo := &repositories.IncidentRepository{Db: db}
	incidentService := &services.IncidentService{Repo: incidentRepo}
	incidentHandler := &handlers.IncidentHandler{Service: incidentService}

	educationRepo := &repositories.EducationRepository{Db: db}
	educationService := &services.EducationService{Repo: educationRepo}
	educationHandler := &handlers.EducationHandler{Service: educationService}

	policeRepo := &repositories.PoliceDepartmentRepository{Db: db}
	policeService := &services.PoliceDepartmentService{Repo: policeRepo}
	policeHandler := &handlers.PoliceDepartmentHandler{Service: policeService}

	emergencyRepo := &repositories.EmergencyRepository{Db: db}
	emergencyService := &services.EmergencyService{Repo: emergencyRepo}
	emergencyHandler := &handlers.EmergencyHandler{
		Service:       emergencyService,
		PoliceService: policeService,
	}

	newsRepo := &repositories.NewsRepository{Db: db}
	newsService := &services.NewsService{Repo: newsRepo}
	newsHandler := &handlers.NewsHandler{Service: newsService}

	messageRepo := &repositories.MessageRepository{Db: db}
	messageService := &services.MessageService{Repo: messageRepo}
	messageHandler := &handlers.MessageHandler{Service: messageService}

	alertRepo := &repositories.AlertRepository{Db: db}
	alertService := &services.AlertService{Repo: alertRepo}
	alertHandler := &handlers.AlertHandler{Service: alertService}

	return &application{
		db:               db,
		errorLog:         errorLog,
		infoLog:          infoLog,
		wsManager:        NewWebSocketManager(),
		userHandler:      userHandler,
		incidentHandler:  incidentHandler,
		educationHandler: educationHandler,
		emergencyHandler: emergencyHandler,
		newsHandler:      newsHandler,
		messageHandler:   messageHandler,
		fcmHandler:       fcmHandler,
		policeHandler:    policeHandler,
		alertHandler:     alertHandler,
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
