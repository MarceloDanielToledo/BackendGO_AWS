package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/jwt"
	"backendgo_aws/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.ResponseAPI {
	var user models.User
	var response models.ResponseAPI
	response.StatusCode = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		response.Message = "Invalid credentials: " + err.Error()
		return response
	}
	if len(user.Email) == 0 {
		response.Message = "Email is required"
		return response
	}
	if len(user.Password) == 0 {
		response.Message = "Password is required"
		return response
	}

	userDB, exist := bd.Login(user.Email, user.Password)
	if !exist {
		response.Message = "Invalid credentials"
		return response
	}

	jwtKey, err := jwt.GenerateJWT(ctx, userDB)
	if err != nil {
		response.Message = "Error generating JWT: " + err.Error()
		return response
	}

	resp := models.LoginResponse{
		Token:   jwtKey,
		Message: "Login success",
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		response.Message = "Error generating JWT: " + err2.Error()
		return response
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}
	cookieString := cookie.String()

	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
		Body: string(token),
	}

	response.StatusCode = 200
	response.Message = string(token)
	response.CustomResp = &res
	return response

}
