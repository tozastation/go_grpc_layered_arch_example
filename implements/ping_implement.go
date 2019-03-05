package implements

import (
	"context"
	"github.com/sirupsen/logrus"
	rpc_ping "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/ping"
)

// IPingImplement is ...
type IPingImplement interface {
}

type pingImplement struct {
	Logger *logrus.Logger
}

// NewPingImplement is ...
func NewPingImplement(l *logrus.Logger) IPingImplement {
	return &pingImplement{l}
}

func (imp *pingImplement) Pong(ctx context.Context, p *rpc_ping.Empty) (*rpc_ping.Pong, error) {
	imp.Logger.Infoln("[START] PongRPC is Called from Client")
	res := rpc_ping.Pong{}
	res.Reply = "Pong"
	imp.Logger.Infoln("[END] PongRPC is Called from Client")
	return &res, nil
}
