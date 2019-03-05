package repository

import (
	"context"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/db"
	rpc_user "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/user"
)

// IUserRepository is ...
type IUserRepository interface {
	FindUserByUserToken(ctx context.Context, token string) (*rpc_user.GetUser, error)
	CreateUser(user *db.User) (string, error)
	Login(uID string, password []byte) (string, error)
}
