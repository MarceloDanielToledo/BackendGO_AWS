package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func AddRelation(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
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

	status, err := bd.AddRelation(t)
	if err != nil {
		response.Message = "Error inserting relation: " + err.Error()
		response.StatusCode = 500
		return response
	}

	if !status {
		response.Message = "Error inserting relation."
		response.StatusCode = 500
		return response
	}

	response.StatusCode = 200
	response.Message = "Relation created successfully"
	return response
}
