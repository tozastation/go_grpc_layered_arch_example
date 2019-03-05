package service

import (
	"context"
	irepo "github.com/tozastation/gRPC-Training-Golang/domain/repository"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/remote"
	rpc_weather "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/weather"
)

// IWeatherService is ...
type IWeatherService interface {
	GetWeather(ctx context.Context, cityName string) (*rpc_weather.Weather, error)
}

type weatherService struct {
	irepo.IWeatherRepository
}

// NewWeatherService is ...
func NewWeatherService(repo irepo.IWeatherRepository) IWeatherService {
	return &weatherService{repo}
}

func (srv *weatherService) GetWeather(ctx context.Context, cityName string) (*rpc_weather.Weather, error) {
	dbWeathers, err := srv.IWeatherRepository.FindCurrentWeatherByCityName(ctx, cityName)
	if err != nil {
		return nil, err
	}
	return dbWeatherToWeather(dbWeathers), nil
}

func dbWeatherToWeather(weather *remote.OpenWeather) *rpc_weather.Weather {
	return &rpc_weather.Weather{
		CityName:    weather.Name,
		TempMax:     weather.Main.TempMax,
		TempMin:     weather.Main.TempMin,
		Wind:        weather.Wind.Speed,
		Type:        weather.Weather[0].Main,
		Description: weather.Weather[0].Description,
	}
}
