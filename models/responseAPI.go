package models

import "github.com/aws/aws-lambda-go/events"

type ResponseAPI struct {
	StatusCode int
	Message    string
	CustomResp *events.APIGatewayProxyResponse
}
