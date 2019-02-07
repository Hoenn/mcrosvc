package server

import (
	"context"

	"github.com/hoenn/mcrosvc/proto"
	"github.com/hoenn/mcrosvc/udb/pkg/db"
)

type UDBServer struct {
	DB db.UserDB
}

func (s *UDBServer) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	u, err := s.DB.GetUser(ctx, in.GetUserNum())
	if err != nil {
		return nil, err
	}
	return &proto.GetUserResponse{
		User: u,
	}, nil

}
func (s *UDBServer) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	u := in.GetUser()
	id, err := s.DB.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserResponse{
		User: &proto.User{
			UserNum: int32(id),
			Age:     u.Age,
			Name:    u.Name,
		},
	}, nil
}
func (s *UDBServer) DeleteUser(ctx context.Context, in *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := s.DB.DeleteUser(ctx, in.GetUserNum())
	if err != nil {
		return nil, err
	}
	return &proto.DeleteUserResponse{}, nil
}
