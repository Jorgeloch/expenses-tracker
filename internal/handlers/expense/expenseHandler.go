package expenseHandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	expenseDTO "github.com/jorgeloch/expenses-tracker/internal/dto/expense"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
)

type Handler struct {
	Service  *service.Service
	Validate *validator.Validate
}

func Init(s *service.Service, v *validator.Validate) *Handler {
	return &Handler{
		Service:  s,
		Validate: v,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ownerID := r.URL.Query().Get("owner_id")
	expenses, err := h.Service.ExpenseService.GetAll(ownerID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expenses)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get the owner id
	ownerID := r.URL.Query().Get("owner_id")

	// get the id from the url
	params := mux.Vars(r)
	// convert the id to int
	id := params["id"]

	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method GetownerByID
	expense, err := h.Service.ExpenseService.GetByID(ownerID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the owner in a json format
	json.NewEncoder(w).Encode(expense)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ownerID := r.URL.Query().Get("owner_id")
	// get the owner from the request body
	var dto expenseDTO.CreateExpenseDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the owner
	err = h.Validate.Struct(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the service method createowner
	location, err := h.Service.ExpenseService.Create(ownerID, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the location of the owner in a json
	json.NewEncoder(w).Encode(location)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ownerID := r.URL.Query().Get("owner_id")

	// get the id from the url
	params := mux.Vars(r)

	id := params["id"]
	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get the owner from the request body
	var dto expenseDTO.UpdateExpenseDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate the owner
	err = h.Validate.Struct(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method updateowner
	expense, err := h.Service.ExpenseService.Update(ownerID, id, dto)
	// return the updated owner in a json format
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(expense)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ownerID := r.URL.Query().Get("owner_id")

	// get the id from the url
	params := mux.Vars(r)

	id := params["id"]
	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method deleteowner
	err := h.Service.ExpenseService.Delete(ownerID, id)
	// return a message in a json format
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
