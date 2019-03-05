package repository

import (
	"context"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/remote"
)

// IWeatherRepository is ...
type IWeatherRepository interface {
	FindCurrentWeatherByCityName(ctx context.Context, cityName string) (*remote.OpenWeather, error)
}
