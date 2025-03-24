package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"context"
	"encoding/json"
	"time"
)

func AddTweet(ctx context.Context, claim models.Claim) models.ResponseAPI {
	var message models.Tweet
	var response models.ResponseAPI
	response.StatusCode = 400
	idUser := claim.ID.Hex()
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		response.Message = "Error unmarshalling data: " + err.Error()
		return response
	}
	if len(message.Message) == 0 {
		response.Message = "Message is required"
		return response
	}

	record := models.AddTweet{
		UserId:  idUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(record)
	if err != nil {
		response.Message = "Error inserting tweet: " + err.Error()
	}
	if !status {
		response.Message = "Error inserting tweet"
	}
	response.StatusCode = 201
	response.Message = "Tweet created successfully"
	return response
}
