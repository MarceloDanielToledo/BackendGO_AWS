package jwt

import (
	"backendgo_aws/models"
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var UserId string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	key := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("invalid token format")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		//Check in the database
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("invalid token")
	}
	return &claims, true, claims.Email, nil

}
