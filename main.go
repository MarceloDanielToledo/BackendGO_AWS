package main

import (
	"backendgo_aws/awsgo"
	"backendgo_aws/bd"
	"backendgo_aws/handlers"
	"backendgo_aws/models"
	"backendgo_aws/secretmanager"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecuteLambda)
}

func EjecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()
	if !ValidateParameters() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error: Missing parameters",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error in get Secret: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	fmt.Println("user", SecretModel.Usernname)
	fmt.Println("password", SecretModel.Password)
	fmt.Println("host", SecretModel.Host)
	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Usernname)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))
	fmt.Println("TestDB")
	// Test DB connection
	err = bd.ConnectDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error in ConnectDB: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	fmt.Println("Gateway")
	fmt.Println("path", awsgo.Ctx.Value(models.Key("path")))
	fmt.Println("method", awsgo.Ctx.Value(models.Key("method")))
	fmt.Println("user", awsgo.Ctx.Value(models.Key("user")))
	fmt.Println("password", awsgo.Ctx.Value(models.Key("password")))
	fmt.Println("host", awsgo.Ctx.Value(models.Key("host")))
	fmt.Println("database", awsgo.Ctx.Value(models.Key("database")))
	fmt.Println("jwtSign", awsgo.Ctx.Value(models.Key("jwtSign")))
	fmt.Println("body", awsgo.Ctx.Value(models.Key("body")))
	fmt.Println("bucketName", awsgo.Ctx.Value(models.Key("bucketName")))

	fmt.Println("Gateway")
	apiResponse := handlers.Handlers(awsgo.Ctx, request)
	if apiResponse.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: apiResponse.StatusCode,
			Body:       apiResponse.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return apiResponse.CustomResp, nil
	}

}

func ValidateParameters() bool {
	_, getParameter := os.LookupEnv("SecretName")
	if !getParameter {
		return false
	}
	_, getParameter = os.LookupEnv("BucketName")
	if !getParameter {
		return false
	}
	_, getParameter = os.LookupEnv("UrlPrefix")
	return getParameter
}
