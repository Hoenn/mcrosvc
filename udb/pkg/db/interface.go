package db

import (
	"context"

	"github.com/hoenn/mcrosvc/proto"
)

type UserDB interface {
	CreateUser(context.Context, *proto.User) (int64, error)
	GetUser(context.Context, int32) (*proto.User, error)
	DeleteUser(context.Context, int32) error
}
