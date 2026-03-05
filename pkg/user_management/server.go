package user_management

import (
	"context"

	pb "github.com/ai-slop-code/pictago/internal/proto/v1/grpc/gateway"
	"github.com/ai-slop-code/pictago/internal/user"
	"github.com/ai-slop-code/pictago/internal/utils"
)

type server struct {
	pb.UnimplementedUserManagementServer
	svcV1 user.ServiceV1
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserInfoResponse, error) {
	usr, err := s.svcV1.GetUserByID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResponse{
		Username: usr.Username,
		Id:       usr.ID.String(),
	}, nil
}

func (s *server) GetUsers(ctx context.Context, _ *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	res, err := s.svcV1.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	userInfos := utils.Map(res, func(t *user.Model) *pb.GetUserInfoResponse {
		return &pb.GetUserInfoResponse{
			Username: t.Username,
			Id:       t.ID.String(),
		}
	})
	return &pb.GetUsersResponse{
		Users: userInfos,
	}, nil
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
