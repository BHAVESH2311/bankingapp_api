package passbookservice

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	id                    string
	Timestamp             time.Time
	SenderId              string
	ReceiverId            string
	SenderAccNumber   string
	ReceivedAccNumber string
	Amount                uint
	TransactionType       string
}

var entry = []*Entry{}

func CreateLog(senderId, receiverId, senderAccNumber, receivedAccNumber string, amount uint, transactionType string) *Entry {
	log := &Entry{
		id:                    uuid.NewString(),
		Timestamp:             time.Now(),
		SenderId:              senderId,
		ReceiverId:            receiverId,
		SenderAccNumber:   senderAccNumber,
		ReceivedAccNumber: receivedAccNumber,
		Amount:                amount,
		TransactionType:       transactionType,
	}

	entry = append(entry, log)

	return log
}

func GetTimeStamp(log *Entry) time.Time {
	return log.Timestamp
}

func FetchPassbook(passbook []*Entry, fromDate, toDate interface{}) ([]*Entry, error) {
    var from, to time.Time
    var err error

    if fromDate != nil {
        fromStr, ok := fromDate.(string)
        if !ok {
            return nil, fmt.Errorf("Invalid fromDate format")
        }

        from, err = time.Parse("2006-01-02", fromStr)
        if err != nil {
            fmt.Println(err)
            panic("Error occurred in Date")
        }
    }

    if toDate != nil {
        toStr, ok := toDate.(string)
        if !ok {
            return nil, fmt.Errorf("Invalid toDate format")
        }

        to, err = time.Parse("2006-01-02", toStr)
        if err != nil {
            fmt.Println(err)
            panic("Error occurred in Date")
        }
    }

    if fromDate == nil && toDate == nil {
        return passbook, nil
    }

    if fromDate == nil {
        var passbookEntry []*Entry
        for _, entry := range passbook {
            if to.Before(GetTimeStamp(entry)) {
                passbookEntry = append(passbookEntry, entry)
            }
        }

        return passbookEntry, nil
    }

    if toDate == nil {
        var passbookEntry []*Entry
        for _, entry := range passbook {
            if from.After(GetTimeStamp(entry)) {
                passbookEntry = append(passbookEntry, entry)
            }
        }

        return passbookEntry, nil
    }

    var passbookEntry []*Entry
    for _, entry := range passbook {
        if from.After(GetTimeStamp(entry)) && to.Before(GetTimeStamp(entry)) {
            passbookEntry = append(passbookEntry, entry)
        }
    }

    return passbookEntry, nil
}


