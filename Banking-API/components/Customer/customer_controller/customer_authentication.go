package usercontroller

import (
	customerservice "bankingapp_api/components/Customer/customer_service"
	auth "bankingapp_api/middlewares/Authentication"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Input struct {
	UserName string
	Password string
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user *customerservice.Customer

	json.NewDecoder(r.Body).Decode(&user)

	newCustomer := customerservice.CreateCustomer(user.Name, user.Password)

	var claims = &auth.Claims{
		ID:      newCustomer.Id,
		Name:    newCustomer.Name,
		IsAdmin: false,
	}

	token, err := auth.SignJWT(*claims)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 10),
		Secure:  true,
	})

	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User has been created Successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		fmt.Println(err)
	}

	customer, err := customerservice.GetCustomerByName(input.UserName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid Credentials"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
		return
	}

	var claims = &auth.Claims{
		ID:      customer.Id,
		Name:    customer.Name,
		IsAdmin: false,
	}

	token, err := auth.SignJWT(*claims)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(" Internal Server Error"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "user-authentication",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 5),
		Secure:  true,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Log In Successful"))
}

