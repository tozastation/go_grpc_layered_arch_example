package implements

import (
	"context"
	"github.com/sirupsen/logrus"
	isrv "github.com/tozastation/go-grpc-ddd-example/domain/service"
	rpc_weather "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/weather"
)

// IWeatherImplement is ...
type IWeatherImplement interface {
	Get(ctx context.Context, p *rpc_weather.GetRequest) (*rpc_weather.GetResponse, error)
}

type weatherImplement struct {
	isrv.IWeatherService
	*logrus.Logger
}

// NewWeatherImplement is ...
func NewWeatherImplement(repo isrv.IWeatherService, logger *logrus.Logger) IWeatherImplement {
	return &weatherImplement{repo, logger}
}

func (imp *weatherImplement) Get(ctx context.Context, p *rpc_weather.GetRequest) (*rpc_weather.GetResponse, error) {
	imp.Logger.Infoln("[START] GetRPC is Called from Client")
	cityName := p.GetCityName()
	imp.Logger.Infoln("[INPUT] CityName=" + cityName)
	weather, err := imp.IWeatherService.GetWeather(ctx, cityName)
	if err != nil {
		return nil, err
	}
	res := rpc_weather.GetResponse{
		Weather: weather,
	}
	imp.Logger.Infoln("[END] GetRPC is Called from Client")
	return &res, nil
}
