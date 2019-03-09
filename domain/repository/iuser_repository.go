package repository

import (
	"context"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/db"
)

// IUserRepository is ...
type IUserRepository interface {
	FindUserByUserToken(ctx context.Context, token string) (*db.User, error)
	CreateUser(user *db.User) (string, error)
	Login(uID, password string) (string, error)
}
