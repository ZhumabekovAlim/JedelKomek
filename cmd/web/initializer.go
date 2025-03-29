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
	errorLog        *log.Logger
	infoLog         *log.Logger
	wsManager       *WebSocketManager
	incidentHandler *handlers.IncidentHandler
}

func initializeApp(db *sql.DB, errorLog, infoLog *log.Logger) *application {

	userRepo := &repositories.UserRepository{Db: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	incidentRepo := &repositories.IncidentRepository{Db: db}
	incidentService := &services.IncidentService{Repo: incidentRepo}
	incidentHandler := &handlers.IncidentHandler{Service: incidentService}

	return &application{
		errorLog:        errorLog,
		infoLog:         infoLog,
		userHandler:     userHandler,
		incidentHandler: incidentHandler,
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
