package bd

import (
	"backendgo_aws/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRelation(t models.Relation) bool {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relations")

	condition := bson.M{
		"userid":         t.UserId,
		"userrelationid": t.UserRelationId,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	return err == nil
}
