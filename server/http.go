package server

import (
	"net/http"

	"authapi/api/user"
	"github.com/gorilla/mux"
)

func ServeHTTP(port string, userHandler user.UserHandler) {
	router := mux.NewRouter()
	router.Path("/register").Methods("POST").HandlerFunc(userHandler.RegisterUser)
	http.ListenAndServe(":"+port, router)
}
