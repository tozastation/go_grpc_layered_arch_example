package di

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/tozastation/go-grpc-ddd-example/domain/service"
	"github.com/tozastation/go-grpc-ddd-example/implements"
	"github.com/tozastation/go-grpc-ddd-example/infrastructure/persistence/mssql"
	custom_error "github.com/tozastation/go-grpc-ddd-example/interfaces/error"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/handler"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/middleware"
	"google.golang.org/grpc"
	"os"
	"time"
)

var logger = logrus.New()
var repoError = custom_error.NewRepositoryError()

// InitializeUser is ...
func InitializeUser() implements.IUserImplement {
	repo := mssql.NewUserRepository(handler.DB, repoError)
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
	fmt.Println("[Done] Initialize gRPC Server")
	return server
}
