package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jorgeloch/expenses-tracker/internal/handlers"
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
)

func main() {
	mux := mux.NewRouter()

	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("impossible to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	// Initialize the Repositories, services and handlers
	repository := repository.Init(db)
	service := service.Init(repository)
	handler := handlers.Init(service)

	mux.HandleFunc("/owner", handler.OwnerHandler.GetAll).Methods("GET")
	mux.HandleFunc("/owner/{id:[0-9]+}", handler.OwnerHandler.GetByID).Methods("GET")
	mux.HandleFunc("/owner", handler.OwnerHandler.Create).Methods("POST")
	mux.HandleFunc("/owner/{id:[0-9]+}", handler.OwnerHandler.Update).Methods("PATCH")
	mux.HandleFunc("/owner/{id:[0-9]+}", handler.OwnerHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
