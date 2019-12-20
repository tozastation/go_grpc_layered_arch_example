package di

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/tozastation/go-grpc-ddd-example/implements"
	"google.golang.org/grpc"
	"os"
	"time"
)

var logger = logrus.New()

// InitializeUser is ...
func InitializeUser() implements.IUserImplement {
	repo := mssql.NewUserRepository(handler.DB, repoError, timeCount)
	srv := service.NewUserService(repo)
	imp := implements.NewUserImplement(srv, logger)
	fmt.Println("[Done] User Injection")
	return imp
}

// InitializeWeather is ...
func InitializeWeather() implements.IWeatherImplement {
	repo := mssql.NewWeatherRepository()
	srv := service.NewWeatherService(repo)
	imp := implements.NewWeatherImplement(srv, logger)
	fmt.Println("[Done] Weather Injection")
	return imp
}

// InitializePing is ...
func InitializePing() implements.IPingImplement {
	imp := implements.NewPingImplement(logger)
	fmt.Println("[Done] Ping Injection")
	return imp
}

// InitializeServer is ...
func InitializeServer() *grpc.Server {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})

	opts := []grpclogrus.Option{
		grpclogrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpclogrus.ReplaceGrpcLogger(logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpcauth.UnaryServerInterceptor(middleware.AuthFunc),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpclogrus.UnaryServerInterceptor(logger, opts...),
		)),
	)
	fmt.Println("[Done] Initialize gRPC Server")
	return server
}
