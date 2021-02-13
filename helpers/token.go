package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("trongnv")

func CreateToken(username string, id string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
