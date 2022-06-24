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

func ConfirmToken(token string, mapClaims jwt.MapClaims) (map[string]interface{}, error) {

	maps := make(map[string]interface{})

	Token, err := jwt.ParseWithClaims(token, mapClaims, func(token *jwt.Token) (interface{}, error) {

		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return maps, err
	}

	//将token中的内容存入parmMap
	claim := Token.Claims.(jwt.MapClaims)
	var parmMap map[string]interface{}
	parmMap = make(map[string]interface{})
	for key, val := range claim {
		parmMap[key] = val
	}

	return claim, err
}
