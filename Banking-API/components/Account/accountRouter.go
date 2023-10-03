package account

import (
	account_controller "bankingapp_api/components/Account/account_controller"

	"github.com/gorilla/mux"
)

func AccountRouter(router *mux.Router) *mux.Router {
	accountRouter := router.PathPrefix("/user/{uId}/bank/{bId}/account").Subrouter()

	accountRouter.HandleFunc("", account_controller.CreateAccount).Methods("POST")
	accountRouter.HandleFunc("/{accountId}", account_controller.GetAccountByID).Methods("GET")
	accountRouter.HandleFunc("/{accountId}", account_controller.UpdateAccount).Methods("POST")

	accountRouter.HandleFunc("/{accountId}/withdraw", account_controller.Withraw).Methods("POST")
	accountRouter.HandleFunc("/{accountId}/deposit", account_controller.Deposit).Methods("POST")

	return accountRouter
}
