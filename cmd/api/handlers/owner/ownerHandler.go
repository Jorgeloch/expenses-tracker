package ownerHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeloch/expenses-tracker/cmd/api/models/owner"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
)

type Handler struct {
	Service *service.Service
}

func Init(s *service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (u *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.Service.OwnerService.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (u *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	params := mux.Vars(r)
	// convert the id to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service method GetUserByID
	user, err := u.Service.OwnerService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the user in a json format
	json.NewEncoder(w).Encode(user)
}

func (u *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// get the user from the request body
	var user ownerModel.Owner
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate the user
	// call the service method createUser
	_, err = u.Service.OwnerService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the location of the user in a json
	json.NewEncoder(w).Encode(user)
}

func (u *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	// get the user from the request body
	// validate the user
	// call the service method updateUser
	// return the updated user in a json format
}

func (u *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	// call the service method deleteUser
	// return a message in a json format
}
