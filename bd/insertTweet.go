package bd

import (
	"backendgo_aws/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.AddTweet) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	register := bson.M{
		"userid":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), true, nil
}
