package debtorHandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	debtorDTO "github.com/jorgeloch/expenses-tracker/internal/dto/debtor"
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
	// get the owner id
	ownerID := context.Get(r, "owner_id").(string)

	debtors, err := h.Service.DebtorService.GetAll(ownerID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(debtors)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get the owner id
	ownerID := context.Get(r, "owner_id").(string)

	// get the id from the url
	params := mux.Vars(r)
	// convert the id to int
	id := params["id"]

	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method GetownerByID
	debtor, err := h.Service.DebtorService.GetByID(ownerID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the owner in a json format
	json.NewEncoder(w).Encode(debtor)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// get the owner id
	ownerID := context.Get(r, "owner_id").(string)

	// get the owner from the request body
	var dto debtorDTO.CreateDebtorDTO

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
	location, err := h.Service.DebtorService.Create(ownerID, dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the location of the owner in a json
	json.NewEncoder(w).Encode(location)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	// get the owner id
	ownerID := context.Get(r, "owner_id").(string)

	// get the id from the url
	params := mux.Vars(r)

	id := params["id"]
	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get the owner from the request body
	var dto debtorDTO.UpdateDebtorDTO

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
	debtor, err := h.Service.DebtorService.Update(ownerID, id, dto)
	// return the updated owner in a json format
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(debtor)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// get the owner id
	ownerID := context.Get(r, "owner_id").(string)

	// get the id from the url
	params := mux.Vars(r)

	id := params["id"]
	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method deleteowner
	err := h.Service.DebtorService.Delete(ownerID, id)
	// return a message in a json format
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
