package user_management

import (
	"context"

	pb "github.com/ai-slop-code/pictago/internal/proto/v1/grpc/gateway"
	"github.com/ai-slop-code/pictago/internal/user"
)

type server struct {
	pb.UnimplementedUserManagementServer
	svcV1 user.ServiceV1
}

func (s *server) GetUsers(context.Context, *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	res := pb.GetUsersResponse{}
	return &res, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	usr, err := s.svcV1.CreateUser(ctx, user.Model{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id:       usr.ID.String(),
		Username: usr.Username,
		Email:    usr.Email,
	}, nil
}

func (s *server) DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res := pb.DeleteUserResponse{}
	return &res, nil
}

var _ pb.UserManagementServer = &server{}

func NewUserManagementServer(svc user.ServiceV1) pb.UserManagementServer {
	return &server{svcV1: svc}
}
