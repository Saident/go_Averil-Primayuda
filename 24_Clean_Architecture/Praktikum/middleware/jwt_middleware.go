package middleware

import (
	"clean_architecture/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
