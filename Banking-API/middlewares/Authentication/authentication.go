package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID      string
	Name    string
	IsAdmin bool
	jwt.StandardClaims
}

var secretKeyJWT = []byte("Secret JWT")

func SignJWT(claims Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := tokenObj.SignedString(secretKeyJWT)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Verify(token string) (*Claims, error) {
	var userClaim = &Claims{}

	tokenObj, err := jwt.ParseWithClaims(token, userClaim, func(t *jwt.Token) (interface{}, error) {
		return secretKeyJWT, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("unauthorized error")
		}

		return nil, errors.New("status bad request")
	}

	if !tokenObj.Valid {
		return nil, errors.New("token invalid")
	}

	return userClaim, nil
}

func AuthenticationHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		token := r.Header.Values("auth")

		if len(token) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token is not Present"))
			return
		}

		tokenStr := token[0]
		_, err := Verify(tokenStr)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Only Admins can Access"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func IsAdmin(function func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		token := r.Header.Values("auth")

		if len(token) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(" Token is not Present"))
			return
		}

		tokenStr := token[0]
		user, err := Verify(tokenStr)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Only Admins can Access"))
			return
		}

		if !user.IsAdmin {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Only Admins can Access"})
			return
		}

		function(w, r) 
	}
}
