package bank

import (
	bank_controller "bankingapp_api/components/Bank/bank_controller"
	auth "bankingapp_api/middlewares/Authentication"

	"github.com/gorilla/mux"
)

func BankRouter(router *mux.Router) *mux.Router {
	bankRouter := router.PathPrefix("/user/{uId}/bank").Subrouter()

	bankRouter.HandleFunc("", bank_controller.FetchAllBanks).Methods("GET")
	bankRouter.HandleFunc("", bank_controller.CreateBank).Methods("POST")
	bankRouter.HandleFunc("/{bId}", auth.IsAdmin(bank_controller.UpdateBank)).Methods("PUT")

	bankRouter.HandleFunc("/{bId}/accounts", bank_controller.FetchAllAccounts).Methods("GET")

	return bankRouter
}
