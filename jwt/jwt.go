package jwt

import (
	"backendgo_aws/models"
	"context"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, t models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	key := []byte(jwtSign)
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"birthdate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.WebSite,
		"avatar":    t.Avatar,
		"banner":    t.Banner,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
