package bd

import (
	"backendgo_aws/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(id string, page int64, search string, typeUser string) ([]*models.User, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var result []*models.User
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetSkip((page - 1) * 20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return result, false
	}

	var include bool
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			fmt.Println("Decode = " + err.Error())
			return result, false
		}

		var relation models.Relation
		relation.UserId = id
		relation.UserRelationId = user.ID.Hex()
		include = false

		findRecord := GetRelation(relation)
		if typeUser == "new" && !findRecord {
			include = true
		}
		if typeUser == "follow" && findRecord {
			include = false
		}

		if relation.UserRelationId == id {
			include = false
		}
		if include {
			user.Password = ""
			result = append(result, &user)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println("cur.Err() = " + err.Error())
		return result, false
	}

	cur.Close(ctx)
	return result, true

}
