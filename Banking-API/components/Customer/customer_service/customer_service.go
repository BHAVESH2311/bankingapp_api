package userservice

import (
	accountservice "bankingapp_api/components/Account/account_service"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	Id       string
	Name     string
	Password string
	IsAdmin  bool
	Accounts []*accountservice.Account
}

var customer = []*Customer{}

func CreateAdmin(Name, Password string) *Customer {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(Password), 5)

	user := &Customer{
		Id:       uuid.NewString(),
		Name:     Name,
		Password: string(hashedPass),
		IsAdmin:  true,
		Accounts: []*accountservice.Account{},
	}

	customer = append(customer, user)
	return user
}

func CreateCustomer(Name, Password string) *Customer {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(Password), 5)
	user := &Customer{
		Id:       uuid.NewString(),
		Name:     Name,
		Password: string(hashedPass),
		IsAdmin:  false,
	}

	customer = append(customer, user)
	return user
}

func GetAllCustomers() []Customer {
	var allCustomers = []Customer{}
	for _, customer := range customer {
		allCustomers = append(allCustomers, *customer)
	}

	return allCustomers
}

func GetCustomerById(id string) (*Customer, error) {
	for _, customer := range customer {
		if customer.Id == id {
			return customer, nil
		}
	}

	return nil, errors.New("Customer Does Not Exist")
}

func GetCustomerByName(name string) (*Customer, error) {
	for _, customer := range customer {
		if customer.Name == name {
			return customer, nil
		}
	}

	return nil, errors.New("User Does Not Exist")
}

func UpdateCustomerById(body *Customer, user *Customer) {
	if body.Name != "" && body.Name != user.Name {
		user.Name = body.Name
	}
}

func DeleteCustomerByID(id string) {
	var index int = 0
	for i, j := range customer {
		if j.Id == id {
			index = i
			break
		}
	}

	customer = append(customer[:index], customer[index+1:]...)
}

func CreateAccount(user *Customer, account *accountservice.Account) {
	user.Accounts = append(user.Accounts, account)
}
