package bd

import (
	"backendgo_aws/models"
	"context"
)

func DeleteRelation(t models.Relation) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
