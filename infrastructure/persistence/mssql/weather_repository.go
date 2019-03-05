package mssql

import (
	"context"
	"fmt"
	irepo "github.com/tozastation/gRPC-Training-Golang/domain/repository"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/remote"
	"io/ioutil"
	"net/http"
	"os"
)

// WeatherRepository is
type WeatherRepository struct {
}

// NewWeatherRepository is ...
func NewWeatherRepository() irepo.IWeatherRepository {
	return &WeatherRepository{}
}

// FindCurrentWeatherByCityName is ...
func (repo *WeatherRepository) FindCurrentWeatherByCityName(ctx context.Context, cityName string) (*remote.OpenWeather, error) {
	baseURL := os.Getenv("OPENWEATHER_URL")
	credential := os.Getenv("OPENWEATHER_CREDENTIAL")
	param := "?q=" + cityName
	res, err := http.Get(baseURL + param + credential)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	openWeather, err := remote.UnmarshalOpenWeather(body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	return &openWeather, nil
}
