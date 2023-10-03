package user

import (
	customercontroller "bankingapp_api/components/Customer/customer_controller"

	"github.com/gorilla/mux"
)

func AuthRouter(router *mux.Router) *mux.Router {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", customercontroller.Login).Methods("POST")
	authRouter.HandleFunc("/register", customercontroller.Register).Methods("POST")

	return authRouter
}
