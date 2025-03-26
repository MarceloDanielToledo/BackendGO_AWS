package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"

	"github.com/aws/aws-lambda-go/events"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		response.Message = "You must send the ID parameter"
		return response
	}

	err := bd.DeleteTweet(ID, claim.ID.Hex())
	if err != nil {
		response.Message = "Error deleting the tweet: " + err.Error()
		return response
	}
	response.StatusCode = 200
	response.Message = "Tweet deleted successfully"
	return response
}
