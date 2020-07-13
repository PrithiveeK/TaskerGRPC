package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client contains the db connection
var Client *mongo.Client

var err error

//DBConnection function connects to the database and returns the client
func DBConnection() {
	ctx, cancel := GetCtx()
	defer cancel()
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err = Client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
}

//GetCtx returns a context with 5 second timeout
func GetCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
