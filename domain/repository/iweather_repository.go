package repository

import (
	"context"
	"github.com/tozastation/go-grpc-ddd-example/infrastructure/persistence/model/remote"
)

// IWeatherRepository is ...
type IWeatherRepository interface {
	FindCurrentWeatherByCityName(ctx context.Context, cityName string) (*remote.OpenWeather, error)
}
