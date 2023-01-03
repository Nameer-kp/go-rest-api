package http

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized ", http.StatusUnauthorized)
			return
		}

	}
}
func validateToken(accesToken string) bool {
	var mySigningKey = []byte("impossible")
	token, err := jwt.Parse(accesToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
func CreateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("impossible"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
