package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func GetUserList(request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]

	idUser := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		response.Message = "You must send the page parameter as an integer greater than 0"
		return response
	}

	users, status := bd.GetAllUsers(idUser, int64(pagTemp), search, typeUser)
	if !status {
		response.Message = "Error getting users: " + err.Error()
		return response
	}

	respJson, err := json.Marshal(users)
	if err != nil {
		response.StatusCode = 500
		response.Message = "Error parsing users: " + err.Error()
		return response
	}

	response.Message = string(respJson)
	response.StatusCode = 200
	return response
}
