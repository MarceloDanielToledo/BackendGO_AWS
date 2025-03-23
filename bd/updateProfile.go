package bd

import (
	"backendgo_aws/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfile(u models.User, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		register["lastname"] = u.LastName
	}
	register["birthdate"] = u.BirthDate
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		register["website"] = u.WebSite
	}
	updateString := bson.M{
		"$set": register,
	}
	objId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err :=
		col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		fmt.Println("Error updating profile: " + err.Error())
		return false, err
	}
	return true, nil
}
