package user

import (
	"encoding/json"
	"net/http"

	userDomain "authapi/domain/user"
)

type RegisterUserRequest struct {
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type RegisterUserReponse struct {
	UserID string `json:"userId"`
}

type UserHandler interface {
	RegisterUser(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userSvc userDomain.Service
}

func NewUserHandler(userSvc userDomain.Service) UserHandler {
	return &userHandler{userSvc}
}

func (handler *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse request
	userReq := userDomain.User{}
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}
	err = handler.userSvc.RegisterNewUser(&userReq)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&userReq)
}
