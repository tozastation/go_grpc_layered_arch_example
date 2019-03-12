package implements

import (
	"context"
	"github.com/sirupsen/logrus"
	isrv "github.com/tozastation/go-grpc-ddd-example/domain/service"
	rpc_user "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/user"
)

// IUserImplement is ...
type IUserImplement interface {
	Get(ctx context.Context, p *rpc_user.GetRequest) (*rpc_user.GetResponse, error)
	Login(ctx context.Context, p *rpc_user.LoginRequest) (*rpc_user.LoginResponse, error)
	Post(ctx context.Context, p *rpc_user.PostRequest) (*rpc_user.PostResponse, error)
}

type userImplement struct {
	isrv.IUserService
	*logrus.Logger
}

// NewUserImplement is ...
func NewUserImplement(s isrv.IUserService, l *logrus.Logger) IUserImplement {
	return &userImplement{s, l}
}

func (imp *userImplement) Get(ctx context.Context, p *rpc_user.GetRequest) (*rpc_user.GetResponse, error) {
	imp.Logger.Infoln("[START] GetMyBoughtVegetablesRPC is Called from Client")
	token := p.GetToken()
	user, err := imp.IUserService.GetMe(ctx, token)
	if err != nil {
		return nil, err
	}
	res := rpc_user.GetResponse{
		User: user,
	}
	imp.Logger.Infoln("[END] GetMyBoughtVegetablesRPC is Called from Client")
	return &res, nil
}

func (imp *userImplement) Login(ctx context.Context, p *rpc_user.LoginRequest) (*rpc_user.LoginResponse, error) {
	token, err := imp.IUserService.SignIn(ctx, p.GetUserID(), p.GetPassword())
	if err != nil {
		return nil, err
	}
	res := rpc_user.LoginResponse{
		CityName: token,
	}
	return &res, nil
}

func (imp *userImplement) Post(ctx context.Context, p *rpc_user.PostRequest) (*rpc_user.PostResponse, error) {
	user := p.GetUser()
	token, err := imp.IUserService.SignUp(ctx, user)
	if err != nil {
		return nil, err
	}
	res := rpc_user.PostResponse{
		CityName: token,
	}
	return &res, nil
}
