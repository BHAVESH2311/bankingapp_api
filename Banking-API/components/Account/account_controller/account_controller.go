package accountcontroller

import (
	accountservice "bankingapp_api/components/Account/account_service"
	bankservice "bankingapp_api/components/Bank/bank_service"
	customerservice "bankingapp_api/components/Customer/customer_service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uId := params["uId"]
	bId := params["bId"]

	user, err := customerservice.GetCustomerById(uId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	bank, err := bankservice.GetBankByID(bId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	var container *accountservice.Account
	json.NewDecoder(r.Body).Decode(&container)
	account := accountservice.CreateAccount(bank.Name, container.AccHolderName, container.AccountType)

	bankservice.CreateAccount(bank, account)

	customerservice.CreateAccount(user, account)

	res := map[string]interface{}{
		"message":       "Account has been Created Successfully",
		"accountInfo": account,
	}

	json.NewEncoder(w).Encode(res)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	params := mux.Vars(r)
	uId := params["uId"]
	bId := params["bId"]
	accId := params["accId"]
	_, err := customerservice.GetCustomerById(uId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	_, err = bankservice.GetBankByID(bId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	account, err := accountservice.GetAccountByID(accId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account Does Not Exist"))
		return
	}

	res := map[string]interface{}{
		"account": account,
	}

	w.Header().Set("ContentType", "application/json")


	json.NewEncoder(w).Encode(res)
}



func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uId := params["uId"]
	bId := params["bId"]
	accId := params["bId"]
	_, err := customerservice.GetCustomerById(uId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Does Not Exist"))
		return
	}

	_, err = bankservice.GetBankByID(bId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	account, err := accountservice.GetAccountByID(accId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account Does Not Exist"))
		return
	}

	var body *accountservice.Account
	json.NewDecoder(r.Body).Decode(&body)

	accountservice.UpdateAccount(body, account)

	res := map[string]interface{}{
		"message": "Account Has Been Updated Successfully",
	}

	w.Header().Set("ContentType", "application/json")


	json.NewEncoder(w).Encode(res)
}



//withdraw function to carry out the withdraw operation


func Withraw(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	uId := params["uId"]
	bId := params["bId"]
	accId := params["accId"]
	_, err := customerservice.GetCustomerById(uId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	_, err = bankservice.GetBankByID(bId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	account, err := accountservice.GetAccountByID(accId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account Does Not Exist"))
		return
	}

	var container struct {
		Sum int
	}

	json.NewDecoder(r.Body).Decode(&container)

	accountservice.Withdraw(account, container.Sum)

	json.NewEncoder(w).Encode("Withdraw has been done Successfully")
}


//deposit function to carry out the deposit operation


func Deposit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	uId := params["uId"]
	bId := params["bId"]
	accId := params["accId"]
	_, err := customerservice.GetCustomerById(uId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Customer Does Not Exist"))
		return
	}

	_, err = bankservice.GetBankByID(bId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	account, err := accountservice.GetAccountByID(accId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account Does Not Exist"))
		return
	}

	var container struct {
		Sum int
	}

	json.NewDecoder(r.Body).Decode(&container)
	accountservice.Deposit(account, container.Sum)

	json.NewEncoder(w).Encode("Deposit has been done Successfully")
}
