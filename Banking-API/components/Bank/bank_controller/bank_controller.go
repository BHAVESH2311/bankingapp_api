package bankcontroller

import (
	
	"encoding/json"
	"net/http"
	accountservice "bankingapp_api/components/Account/account_service"
	bankservice "bankingapp_api/components/Bank/bank_service"

	"github.com/gorilla/mux"
)

func CreateBank(w http.ResponseWriter, r *http.Request) {

	var container struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&container)

	var newBank = bankservice.CreateBank(container.Name)

	w.Header().Set("ContentType", "application/json")

	json.NewEncoder(w).Encode(newBank)
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {

	bId := mux.Vars(r)["bId"]

	bank, err := bankservice.GetBankByID(bId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	var container bankservice.Bank
	json.NewDecoder(r.Body).Decode(&container)

	bankservice.UpdateBank(bank, container.Name)

	w.Header().Set("ContentType", "application/json")


	json.NewEncoder(w).Encode(map[string]string{"message": "Bank has been Updated Successfully"})
}



func FetchAllBanks(w http.ResponseWriter, r *http.Request) {

	banks := bankservice.GetAllBanks()

	w.Header().Set("ContentType", "application/json")


	json.NewEncoder(w).Encode(map[string][]bankservice.Bank{"banks": banks})
}



func FetchAllAccounts(w http.ResponseWriter, r *http.Request) {

	bId := mux.Vars(r)["bId"]

	bank, err := bankservice.GetBankByID(bId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Bank Does Not Exist"))
		return
	}

	var Accounts = []accountservice.Account{}

	for _, account := range bank.Accounts {
		Accounts = append(Accounts, *account)
	}

	w.Header().Set("ContentType", "application/json")

	json.NewEncoder(w).Encode(Accounts)
}
