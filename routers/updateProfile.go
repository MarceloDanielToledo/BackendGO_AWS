package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"context"
	"encoding/json"
)

func UpdateProfile(ctx context.Context, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400

	var t models.User

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		response.Message = err.Error()
		return response
	}
	var status bool
	status, err = bd.UpdateProfile(t, claim.ID.Hex())
	if err != nil {
		response.Message = "Error updating profile: " + err.Error()
		return response
	}
	if !status {
		response.Message = "Error updating profile"
		return response
	}

	response.StatusCode = 200
	response.Message = "Profile updated successfully"
	return response
}
