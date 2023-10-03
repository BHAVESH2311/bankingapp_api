package accountservice

import (
	passbookservice "bankingapp_api/components/Log/log_service"
	"errors"

	"github.com/google/uuid"
)

type Account struct {
	Id            string
	AccNum        string
	AccountType   string
	Balance       int
	AccHolderName string
	BankName      string
	IsActive      bool
	Passbook      []*passbookservice.Entry
}

var accounts = []*Account{}

func CreateAccount(bankName, accHolderName, accountType string) *Account {

	accountId := uuid.NewString()
	accountNum := uuid.NewString()

	account := &Account{
		Id:            accountId,
		AccNum:        accountNum,
		AccHolderName: accHolderName,
		Balance:       0,
		BankName:      bankName,
		AccountType:   accountType,
		IsActive:      true,
		Passbook:      []*passbookservice.Entry{},
	}

	accounts = append(accounts, account)
	return account
}

func UpdateAccount(req *Account, account *Account) {
	if req.AccHolderName != "" && req.AccHolderName != account.AccHolderName {
		account.AccHolderName = req.AccHolderName
	}
}

func GetAccountByID(id string) (*Account, error) {
	for _, account := range accounts {
		if account.Id == id {
			return account, nil
		}
	}

	return nil, errors.New("Account does not exist")
}

func Deposit(account *Account, amount int) error {
	if amount < 0 {
		panic("amount must be greater than 0")
	}

	account.Balance += amount

	log := passbookservice.CreateLog(account.Id, account.Id, account.AccNum, account.AccNum, uint(amount), "Deposited")

	account.Passbook = append(account.Passbook, log)
	return nil
}

func Withdraw(account *Account, amount int) error {
	if amount < 0 {
		panic("Amount must be greater than 0")
	}

	if amount > account.Balance {
		panic("Low Balance")
	}

	account.Balance -= amount
	log := passbookservice.CreateLog(account.Id, account.Id, account.AccNum, account.AccNum, uint(amount), "Withdrawal")
	account.Passbook = append(account.Passbook, log)
	return nil
}
