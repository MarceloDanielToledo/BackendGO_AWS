package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func GetFollowerTweets(request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400
	idUser := claim.ID.Hex()
	page := request.QueryStringParameters["page"]
	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		response.Message = "You must send the page parameter as an integer greater than 0"
		return response
	}

	tweets, success := bd.GetFollowersTweets(idUser, pag)
	if !success {
		response.Message = "Error reading the tweet"
		response.StatusCode = 500
		return response
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		response.StatusCode = 500
		response.Message = "Error parsing tweets: " + err.Error()
		return response
	}

	response.StatusCode = 200
	response.Message = string(respJson)
	return response
}
