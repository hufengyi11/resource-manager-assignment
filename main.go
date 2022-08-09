package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	gen "assignment/gen/go"
	impl "assignment/impl"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

var db *mongo.Client
var projectdb *mongo.Collection
var monogoCtx context.Context

// export MONGODB_URI='mongodb+srv://dbUser:dbUserPassword@cluster0.kjucuqb.mongodb.net/?retryWrites=true&w=majority'
func main() {

	opts := []grpc.ServerOption{}
	var grpcServer = grpc.NewServer(opts...)
	var assignmentServer = impl.AssignmentServiceServerImpl{}
	gen.RegisterAssignmentServiceServer(grpcServer, &assignmentServer)

	address := ":8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("could not listen to %s %v !!!!!!", address, err)
	}

	fmt.Println("successfully connect to grpc at port:", address)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Printf("Connect error: %v \n", err)
	}
	defer client.Disconnect(context.Background())

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Printf("Ping Error: %v \n", err)
	}

	// assignmentCollection := client.Database("Assignment").Collection("assignment")

	// assignment := bson.D{{"id", 1}, {"project", 1}, {"people", 1}}
	// result, err := assignmentCollection.InsertOne(context.Background(), assignment)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result.InsertedID)
	log.Fatalln(grpcServer.Serve(lis))

}
