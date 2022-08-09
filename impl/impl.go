package impl

import (
	"context"
	"fmt"
	"log"
	"os"

	gen "assignment/gen/go"

	"google.golang.org/grpc"
	gRPC "google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type AssignmentServiceServerImpl struct {
	gen.UnimplementedAssignmentServiceServer
}

type AssignmentItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID string             `bson:"peoject_id"`
	PeopleID  string             `bson:"people_id"`
}

func mongoNewClient() (*mongo.Client, *mongo.Collection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, nil, err
	}

	resourceManagerDB := client.Database("Assignment")
	assignmentCollection := resourceManagerDB.Collection("assignment")

	return client, assignmentCollection, nil
}

func projectServiceConnection() *grpc.ClientConn {
	serverAddress := "rm-p-ms-t-srv:8080"
	// serverAddress := "host.docker.internal:8080"

	conn, e := gRPC.Dial(serverAddress, gRPC.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		fmt.Printf("grpc connection failed: %v", e)
	}
	defer conn.Close()

	return conn
}

func (U *AssignmentServiceServerImpl) CreateAssignment(ctx context.Context, req *gen.CreateAssignmentReq) (*gen.CreateAssignmentRes, error) {
	client, collection, err := mongoNewClient()
	if err != nil {
		log.Fatalf("mongoNewClient error %v \n", err)
	}
	conn := projectServiceConnection()
	client := gen.NewAssignmentServiceClient(conn)
	assignment := req.GetAssignment()
	// verify if the project exists, and person exists

	res, error := collection.InsertOne(
		ctx, bson.D{
			{Key: "Id", Value: assignment.Id},
			{Key: "ProjectId", Value: assignment.ProjectId},
			{Key: "PeopleId", Value: assignment.PeopleId},
		})

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(res.InsertedID)

	return &gen.CreateAssignmentRes{Assignment: assignment}, nil
}

func (U *AssignmentServiceServerImpl) ReadByPeopleAssignment(context.Context, *gen.ReadAssignmentByPeopleReq) (*gen.ReadAssignmentByPeopleRes, error) {

	return nil, status.Errorf(codes.Unimplemented, "method ReadByPeopleAssignment not implemented")
}

func (U *AssignmentServiceServerImpl) ReadByProjectAssignment(context.Context, *gen.ReadAssignmentByProjectReq) (*gen.ReadAssignmentByProjectRes, error) {

	return nil, status.Errorf(codes.Unimplemented, "method ReadByProjectAssignment not implemented")
}

func (U *AssignmentServiceServerImpl) UpdateAssignment(context.Context, *gen.UpdateAssignmentReq) (*gen.UpdateAssignmentRes, error) {

	return nil, status.Errorf(codes.Unimplemented, "method UpdateAssignment not implemented")
}

func (U *AssignmentServiceServerImpl) DeleteAssignment(context.Context, *gen.DeleteAssignmentReq) (*gen.DeleteAssignmentRes, error) {

	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssignment not implemented")
}
