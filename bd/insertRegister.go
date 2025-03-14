package bd

import (
	"backendgo_aws/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(u models.User) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}
