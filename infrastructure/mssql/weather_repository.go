package mssql

import (
	"context"
	"database/sql"
	"encoding/json"
	irepo "github.com/tozastation/gRPC-Training-Golang/domain/repository"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/remote"
	"io/ioutil"
	"net/http"
	"os"
)

// WeatherRepository is
type WeatherRepository struct {
	DB *sql.DB
}

// NewWeatherRepository is ...
func NewWeatherRepository(Conn *sql.DB) irepo.IWeatherRepository {
	return &WeatherRepository{Conn}
}

// FindCurrentWeatherByCityName is ...
func (repo *WeatherRepository) FindCurrentWeatherByCityName(ctx context.Context, cityName string) (*remote.OpenWeather, error) {
	openWeather := remote.OpenWeather{}
	baseURL := os.Getenv("OPENWEATHER_URL")
	credential := os.Getenv("OPENWEATHER_CREDENTIAL")
	param := "?=" + cityName
	res, err := http.Get(baseURL + param + credential)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &openWeather)
	if err != nil {
		return nil, err
	}
	return &openWeather, nil
}
