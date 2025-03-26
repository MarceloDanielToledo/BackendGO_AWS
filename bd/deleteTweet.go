package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserID string) error {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return err
	}
	return nil
}
