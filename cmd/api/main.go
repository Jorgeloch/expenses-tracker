package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jorgeloch/expenses-tracker/internal/handlers"
	"github.com/jorgeloch/expenses-tracker/internal/initializers"
	"github.com/jorgeloch/expenses-tracker/internal/middlewares"
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
)

func init() {
	initializers.LoadENV()
}

func main() {
	mux := mux.NewRouter()

	db, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("impossible to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	// Initialize the Repositories, services and handlers
	repository := repository.Init(db)
	service := service.Init(repository)
	handler := handlers.Init(service)

	mux.Use(middlewares.Logger)

	mux.HandleFunc("/owner/login", handler.OwnerHandler.Login).Methods("POST")
	mux.HandleFunc("/owner", handler.OwnerHandler.Create).Methods("POST")
	mux.HandleFunc("/owner", handler.OwnerHandler.GetAll).Methods("GET")
	mux.HandleFunc("/owner/{id}", middlewares.ValidateToken(handler.OwnerHandler.GetByID)).Methods("GET")
	mux.HandleFunc("/owner/{id}", middlewares.ValidateToken(handler.OwnerHandler.Update)).Methods("PATCH")
	mux.HandleFunc("/owner/{id}", middlewares.ValidateToken(handler.OwnerHandler.Delete)).Methods("DELETE")

	mux.HandleFunc("/card", middlewares.ValidateToken(handler.CardHandler.GetAll)).Methods("GET")
	mux.HandleFunc("/card/{id}", middlewares.ValidateToken(handler.CardHandler.GetByID)).Methods("GET")
	mux.HandleFunc("/card", middlewares.ValidateToken(handler.CardHandler.Create)).Methods("POST")
	mux.HandleFunc("/card/{id}", middlewares.ValidateToken(handler.CardHandler.Update)).Methods("PATCH")
	mux.HandleFunc("/card/{id}", middlewares.ValidateToken(handler.CardHandler.Delete)).Methods("DELETE")

	mux.HandleFunc("/debtor", middlewares.ValidateToken(handler.DebtorHandler.GetAll)).Methods("GET")
	mux.HandleFunc("/debtor/{id}", middlewares.ValidateToken(handler.DebtorHandler.GetByID)).Methods("GET")
	mux.HandleFunc("/debtor", middlewares.ValidateToken(handler.DebtorHandler.Create)).Methods("POST")
	mux.HandleFunc("/debtor/{id}", middlewares.ValidateToken(handler.DebtorHandler.Update)).Methods("PATCH")
	mux.HandleFunc("/debtor/{id}", middlewares.ValidateToken(handler.DebtorHandler.Delete)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
