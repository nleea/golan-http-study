package user

import (
	auth "api-go/services/auth"
	"api-go/types"
	utils "api-go/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlre struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handlre {
	return &Handlre{
		store: store,
	}
}

func (h *Handlre) RegisterRouter(route *mux.Router) {
	route.HandleFunc("/login", h.HandleLogin).Methods("POST")
	route.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handlre) HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlre) HandleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exist", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.CreateUser(types.RegisterUserPayload{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
