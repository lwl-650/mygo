package util

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("Key of Chery")

func GenerateToken(mapClaims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	return token.SignedString(mySigningKey)

}

func ConfirmToken(token string, mapClaims jwt.MapClaims) (err error) {
	Token, err := jwt.ParseWithClaims(token, mapClaims, func(token *jwt.Token) (interface{}, error) {

		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("++++", Token.Claims, "++++")
	for k, v := range Token.Claims {
		fmt.Println(k, v)
	}
	return err
}
