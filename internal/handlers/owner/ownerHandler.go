package ownerHandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	ownerDTO "github.com/jorgeloch/expenses-tracker/internal/dto/owner"
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
	owners, err := h.Service.OwnerService.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(owners)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var dto ownerDTO.LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validate.Struct(dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := h.Service.OwnerService.Login(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	params := mux.Vars(r)
	// convert the id to int
	id := params["id"]

	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method GetownerByID
	owner, err := h.Service.OwnerService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the owner in a json format
	json.NewEncoder(w).Encode(owner)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// get the owner from the request body
	var dto ownerDTO.CreateOwnerDTO
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
	location, err := h.Service.OwnerService.Create(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the location of the owner in a json
	json.NewEncoder(w).Encode(location)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	params := mux.Vars(r)

	id := params["id"]
	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get the owner from the request body
	var dto ownerDTO.UpdateOwnerDTO

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
	owner, err := h.Service.OwnerService.Update(id, dto)
	// return the updated owner in a json format
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(owner)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	params := mux.Vars(r)

	id := params["id"]
	if err := uuid.Validate(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method deleteowner
	err := h.Service.OwnerService.Delete(id)
	// return a message in a json format
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
