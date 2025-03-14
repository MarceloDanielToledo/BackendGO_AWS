package bd

import (
	"backendgo_aws/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckAlreadyExist(email string) (models.User, bool, string) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")
	condition := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	id := result.ID.Hex()
	if err != nil {
		return result, false, id
	}
	return result, true, id
}
