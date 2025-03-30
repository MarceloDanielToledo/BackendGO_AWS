package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func GetRelation(request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		response.Message = "You must send the ID parameter"
		return response
	}

	var t models.Relation
	t.UserId = claim.ID.Hex()
	t.UserRelationId = ID

	var resp models.RelationResponse

	existRelation := bd.GetRelation(t)
	resp.Status = existRelation

	respJson, err := json.Marshal(existRelation)
	if err != nil {
		response.StatusCode = 500
		response.Message = "Error parsing relation: " + err.Error()
		return response
	}

	response.StatusCode = 200
	response.Message = string(respJson)
	return response
}
