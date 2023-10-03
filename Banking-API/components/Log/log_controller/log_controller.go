package logcontroller
import (
	passbookservice "bankingapp_api/components/Log/log_service"
	"net/http"
	"encoding/json"

)
func FetchPassbookEntries(w http.ResponseWriter, r *http.Request) {
    fromDate := r.FormValue("fromDate")
    toDate := r.FormValue("toDate")

	var entry *passbookservice.Entry

    entries, err := passbookservice.FetchPassbook(entry, fromDate, toDate)

    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    responseJSON, err := json.Marshal(entries)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseJSON)
}

