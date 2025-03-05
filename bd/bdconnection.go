package bd

import (
	"backendgo_aws/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)
	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error in ConnectDB: ", err.Error())
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error in Ping: ", err.Error())
		return err
	}
	fmt.Println("Connected to MongoDB")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func DatabaseConnected() bool {
	err := MongoCN.Ping(context.Background(), nil)
	if err != nil {
		return false
	}
	return true
}
