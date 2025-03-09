package handlers

import (
	"backendgo_aws/jwt"
	"backendgo_aws/models"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ResponseAPI {
	fmt.Println("Processing request: " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var r models.ResponseAPI
	r.StatusCode = 400

	isOk, statusCode, message, claim := validateAuthorization(ctx, request)

	if !isOk {
		r.StatusCode = statusCode
		r.Message = message
		return r

	}

	switch ctx.Value(models.Key("method")).(string) {
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)

		}
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	default:

	}
	r.Message = "Method Invalid"
	return r

}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "OK", models.Claim{}
	}
	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token required", models.Claim{}
	}

	claim, allOk, message, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !allOk {
		if err != nil {
			fmt.Println("Error in token verification " + err.Error())
			return false, 401, "Error in token verification " + err.Error(), models.Claim{}
		}
		fmt.Println("Token invalid " + message)
		return false, 401, "Token invalid " + message, models.Claim{}
	}
	fmt.Println("Token ok")
	return true, 200, "OK", *claim

}
