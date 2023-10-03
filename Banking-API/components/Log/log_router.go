package user

import (
	passbookcontroller "bankingapp_api/components/Log/log_controller"

	"github.com/gorilla/mux"
)

func CustomerRouter(router *mux.Router) *mux.Router {
	logRouter := router.PathPrefix("/user/{uId}/bank/{bId}/account").Subrouter()

	logRouter.HandleFunc("/{accountId}", passbookcontroller.FetchPassbookEntries).Methods("GET")

	return logRouter
}
