package user

import (
	customercontroller "bankingapp_api/components/Customer/customer_controller"

	"github.com/gorilla/mux"
)

func CustomerRouter(router *mux.Router) *mux.Router {
	customerRouter := router.PathPrefix("/customer").Subrouter()

	customerRouter.HandleFunc("", customercontroller.GetAllCustomers).Methods("GET")
	customerRouter.HandleFunc("", customercontroller.CreateCustomer).Methods("POST")
	customerRouter.HandleFunc("/{id}", customercontroller.GetCustomerById).Methods("GET")
	customerRouter.HandleFunc("/{id}", customercontroller.UpdateCustomerById).Methods("PUT")
	customerRouter.HandleFunc("/{id}", customercontroller.DeleteCustomerByID).Methods("DELETE")

	customerRouter.HandleFunc("/{id}/myaccounts", customercontroller.GetAllAccount).Methods("GET")

	return customerRouter
}
