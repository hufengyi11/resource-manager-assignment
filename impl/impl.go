package impl

import (
	"context"

	gen "assignment/gen/go"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
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

func (U *AssignmentServiceServerImpl) CreateAssignment(context.Context, *gen.CreateAssignmentReq) (*gen.CreateAssignmentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssignment not implemented")
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
