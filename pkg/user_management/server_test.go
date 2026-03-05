package user_management

import (
	"context"
	"errors"
	"testing"

	pb "github.com/ai-slop-code/pictago/internal/proto/v1/grpc/gateway"
	"github.com/ai-slop-code/pictago/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockServiceV1 struct {
	usr user.Model
	err error
}

func (m *mockServiceV1) GetAllUsers(context.Context) ([]*user.Model, error) {
	return []*user.Model{&m.usr}, m.err
}

func (m *mockServiceV1) GetUserByID(context.Context, string) (*user.Model, error) {
	return &m.usr, m.err
}

func (m *mockServiceV1) CreateUser(context.Context, user.Model) (*user.Model, error) {
	return &m.usr, m.err
}

var _ user.ServiceV1 = &mockServiceV1{}

func TestServer_CreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		mockSvc := &mockServiceV1{
			usr: user.Model{},
		}
		srv := NewUserManagementServer(mockSvc)
		res, err := srv.CreateUser(t.Context(), &pb.CreateUserRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should return error when user already exists", func(t *testing.T) {
		mockSvc := &mockServiceV1{
			err: errors.New("user already exists"),
		}
		srv := NewUserManagementServer(mockSvc)
		_, err := srv.CreateUser(t.Context(), &pb.CreateUserRequest{})
		require.Error(t, err)
	})
}

func TestServer_GetUserByID(t *testing.T) {
	t.Run("should get user by id", func(t *testing.T) {
		mockSvc := &mockServiceV1{
			usr: user.Model{},
		}
		srv := NewUserManagementServer(mockSvc)
		res, err := srv.GetUser(t.Context(), &pb.GetUserRequest{})
		require.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should return error when user does not exists", func(t *testing.T) {
		mockSvc := &mockServiceV1{
			err: errors.New("user does not exists"),
		}
		srv := NewUserManagementServer(mockSvc)
		_, err := srv.GetUser(t.Context(), &pb.GetUserRequest{})
		require.Error(t, err)
	})
}

func TestServer_GetUsers(t *testing.T) {
	t.Run("should get users", func(t *testing.T) {
		mockSvc := &mockServiceV1{
			usr: user.Model{},
		}
		srv := NewUserManagementServer(mockSvc)
		res, err := srv.GetUsers(t.Context(), &pb.GetUsersRequest{})
		require.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should return error when user does not exists", func(t *testing.T) {
		mockSvc := &mockServiceV1{
			err: errors.New("user does not exists"),
		}
		srv := NewUserManagementServer(mockSvc)
		_, err := srv.GetUsers(t.Context(), &pb.GetUsersRequest{})
		require.Error(t, err)
	})
}
