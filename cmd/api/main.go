package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jorgeloch/expenses-tracker/internal/handlers"
	"github.com/jorgeloch/expenses-tracker/internal/middlewares"
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

	mux.Use(middlewares.Logger)

	mux.HandleFunc("/owner", handler.OwnerHandler.GetAll).Methods("GET")
	mux.HandleFunc("/owner/{id}", handler.OwnerHandler.GetByID).Methods("GET")
	mux.HandleFunc("/owner", handler.OwnerHandler.Create).Methods("POST")
	mux.HandleFunc("/owner/{id}", handler.OwnerHandler.Update).Methods("PATCH")
	mux.HandleFunc("/owner/{id}", handler.OwnerHandler.Delete).Methods("DELETE")

	mux.HandleFunc("/card", handler.CardHandler.GetAll).Methods("GET")
	mux.HandleFunc("/card/{id}", handler.CardHandler.GetByID).Methods("GET")
	mux.HandleFunc("/card", handler.CardHandler.Create).Methods("POST")
	mux.HandleFunc("/card/{id}", handler.CardHandler.Update).Methods("PATCH")
	mux.HandleFunc("/card/{id}", handler.CardHandler.Delete).Methods("DELETE")

	mux.HandleFunc("/debtor", handler.DebtorHandler.GetAll).Methods("GET")
	mux.HandleFunc("/debtor/{id}", handler.DebtorHandler.GetByID).Methods("GET")
	mux.HandleFunc("/debtor", handler.DebtorHandler.Create).Methods("POST")
	mux.HandleFunc("/debtor/{id}", handler.DebtorHandler.Update).Methods("PATCH")
	mux.HandleFunc("/debtor/{id}", handler.DebtorHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
