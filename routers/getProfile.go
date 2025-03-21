package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ViewProfile(request events.APIGatewayProxyRequest) models.ResponseAPI {
	var response = models.ResponseAPI{
		StatusCode: 400,
	}
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		response.Message = "ID is required"
		return response
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		response.Message = "Error getting profile: " + err.Error()
		return response
	}

	responseJson, err := json.Marshal(profile)
	if err != nil {
		response.StatusCode = 500
		response.Message = "Error parsing profile: " + err.Error()
		return response
	}

	response.StatusCode = 200
	response.Message = string(responseJson)
	return response
}
