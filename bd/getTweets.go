package bd

import (
	"backendgo_aws/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(Id string, page int64) ([]*models.TweetResponse, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")
	var results []*models.TweetResponse

	condition := bson.M{
		"userid": Id,
	}
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, findOptions)
	if err != nil {
		return results, false
	}

	for cursor.Next(ctx) {
		var register models.TweetResponse
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
