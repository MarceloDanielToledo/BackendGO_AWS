package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"

	"github.com/aws/aws-lambda-go/events"
)

func DeleteRelation(request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
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

	status, err := bd.DeleteRelation(t)
	if err != nil {
		response.Message = "Error deleting relation: " + err.Error()
		return response
	}

	if !status {
		response.Message = "Error deleting relation"
		return response
	}

	response.StatusCode = 200
	response.Message = "Relation deleted successfully"
	return response
}
