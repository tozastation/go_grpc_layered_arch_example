package di

import (
	"github.com/sirupsen/logrus"
	"github.com/tozastation/gRPC-Training-Golang/domain/service"
	"github.com/tozastation/gRPC-Training-Golang/implements"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/mssql"
	"github.com/tozastation/gRPC-Training-Golang/interfaces/handler"
)

var conn = handler.OpenDBConnection()
var logger = logrus.New()

// InitializeUser is ...
func InitializeUser() implements.IUserImplement {
	repo := mssql.NewUserRepository(conn)
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
