package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func ReadTweet(request events.APIGatewayProxyRequest) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400

	ID := request.QueryStringParameters["id"]
	page := request.QueryStringParameters["page"]
	if len(ID) < 1 {
		response.Message = "You must send the ID parameter"
		return response
	}
	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		response.Message = "You must send the page parameter as an integer greater than 0"
		return response
	}

	tweet, success := bd.GetTweets(ID, int64(pag))
	if !success {
		response.Message = "Error reading the tweet"
		return response
	}
	respoJson, err := json.Marshal(tweet)
	if err != nil {
		response.Message = "Error converting the tweet"
		return response
	}
	response.Message = string(respoJson)
	response.StatusCode = 200
	return response
}
