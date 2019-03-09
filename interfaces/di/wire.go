package di

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/tozastation/gRPC-Training-Golang/domain/service"
	"github.com/tozastation/gRPC-Training-Golang/implements"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/mssql"
	"github.com/tozastation/gRPC-Training-Golang/interfaces/handler"
	"github.com/tozastation/gRPC-Training-Golang/interfaces/middleware"
	"google.golang.org/grpc"
	"os"
	"time"
)

var logger = logrus.New()

// InitializeUser is ...
func InitializeUser() implements.IUserImplement {
	repo := mssql.NewUserRepository(handler.OpenDBConnection())
	srv := service.NewUserService(repo)
	imp := implements.NewUserImplement(srv, logger)
	return imp
}

// InitializeWeather is ...
func InitializeWeather() implements.IWeatherImplement {
	repo := mssql.NewWeatherRepository()
	srv := service.NewWeatherService(repo)
	imp := implements.NewWeatherImplement(srv, logger)
	return imp
}

// InitializePing is ...
func InitializePing() implements.IPingImplement {
	imp := implements.NewPingImplement(logger)
	return imp
}

// InitializeServer is ...
func InitializeServer() *grpc.Server {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})

	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpc_logrus.ReplaceGrpcLogger(logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middleware.AuthFunc),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
		)),
	)
	return server
}
