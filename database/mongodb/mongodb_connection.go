package mongodb

import (
	"api/helpers"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConn() {
	connectionString := helpers.NewDotEnvHelper().GoDotEnvVariable("MONGODB_CONN_STRING")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil{
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
}