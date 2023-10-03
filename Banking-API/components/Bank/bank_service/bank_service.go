package bankservice

import (
	accountservice "bankingapp_api/components/Account/account_service"
	"errors"

	"github.com/google/uuid"
)

type Bank struct {
	Bid       string
	Name      string
	NetWorth  int
	Accounts []*accountservice.Account
}

var banks = []*Bank{}

func CreateBank(name string) *Bank {

	bankId := uuid.NewString()

	bank := &Bank{
		Bid:       bankId,
		Name:     name,
		NetWorth: 0,
		Accounts: []*accountservice.Account{},
	}

	banks = append(banks, bank)
	return bank
}

func GetAllBanks() []Bank {
	BanksContainer := []Bank{}

	for _, bank := range banks {
		BanksContainer = append(BanksContainer, *bank)
	}

	return BanksContainer
}

func UpdateBank(bank *Bank, bankName string) {
	bank.Name = bankName
}

func GetBankByID(bankId string) (*Bank, error) {
	for _, bank := range banks {
		bank.Bid = bankId
		return bank, nil
	}

	return nil, errors.New("Bank not found")
}

func CreateAccount(bank *Bank, account *accountservice.Account) {
	bank.Accounts = append(bank.Accounts, account)
}
