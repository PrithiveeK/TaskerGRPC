package main

import (
	"log"
	"net"

	"../config"
	"../protos"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type taskServer struct{}

//ProjectsColl in not  exported and contains projects collection
var ProjectsColl *mongo.Collection

//TasksColl in not  exported and contains tasks collection
var TasksColl *mongo.Collection

//TeamsColl in not  exported and contains teams collection
var TeamsColl *mongo.Collection

//UsersColl in not  exported and contains users collection
var UsersColl *mongo.Collection

func main() {
	client := config.Client
	ctx, cancel := config.GetCtx()
	taskerDB := client.Database("taskerapp")

	defer client.Disconnect(ctx)
	defer cancel()

	ProjectsColl = taskerDB.Collection("projects")
	TasksColl = taskerDB.Collection("tasks")
	TeamsColl = taskerDB.Collection("teams")
	UsersColl = taskerDB.Collection("users")

	server := grpc.NewServer()
	protos.RegisterTaskerAppServer(server, taskServer{})
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("Server starting...", server.Serve(listener).Error())
}
