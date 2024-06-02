package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) HandleUserRegistration(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleUserLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleUserRegistration).Methods("POST")
}
