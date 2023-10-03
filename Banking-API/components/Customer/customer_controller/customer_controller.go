package usercontroller

import (
	accountservice "bankingapp_api/components/Account/account_service"
	customerservice "bankingapp_api/components/Customer/customer_service"
	validators "bankingapp_api/validators"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAdmin(w http.ResponseWriter, r *http.Request) {

	var customer customerservice.Customer
	json.NewDecoder(r.Body).Decode(&customer)

	if !validators.ValidateName(customer.Name) {
		panic("Enter valid Name")
	}

	if !validators.ValidatePassword(customer.Password) {
		panic("Enter valid Password")
	}

	var newCustomer = customerservice.CreateAdmin(customer.Name, customer.Password)
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(newCustomer)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer customerservice.Customer
	json.NewDecoder(r.Body).Decode(&customer)

	if !validators.ValidateName(customer.Name) {
		panic("Enter valid Name")
	}

	if !validators.ValidatePassword(customer.Password) {
		panic("Enter valid Password")
	}
	var newcustomer = customerservice.CreateCustomer(customer.Name, customer.Password)
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(newcustomer)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers := customerservice.GetAllCustomers()

	w.Header().Set("ContentType", "application/json")


	json.NewEncoder(w).Encode(map[string][]customerservice.Customer{"customers": customers})
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)

	customers, err := customerservice.GetCustomerById(mp["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Does Not Exist"))
		return
	}
	w.Header().Set("ContentType", "application/json")


	json.NewEncoder(w).Encode(customers)
}

func UpdateCustomerById(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)

	customer, err := customerservice.GetCustomerById(mp["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Does Not Exist"))
		return
	}

	var body *customerservice.Customer
	json.NewDecoder(r.Body).Decode(&body)

	customerservice.UpdateCustomerById(body, customer)

	w.Header().Set("ContentType", "application/json")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"Action": "Customer has been Updated Successfully"})
}

func DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)

	_, err := customerservice.GetCustomerById(mp["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Does Not Exist"))
		return
	}

	customerservice.DeleteCustomerByID(mp["id"])
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Action": "Customer has been Deleted Successfully"})
}

func GetAllAccount(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	cId := params["id"]
	customer, err := customerservice.GetCustomerById(cId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Does Not Exist"))
		return
	}

	allAccounts := []accountservice.Account{}

	for _, account := range customer.Accounts {
		allAccounts = append(allAccounts, *account)
	}

	res := map[string]interface{}{
		"accounts": allAccounts,
	}

	w.Header().Set("ContentType", "application/json")

	json.NewEncoder(w).Encode(res)
}
