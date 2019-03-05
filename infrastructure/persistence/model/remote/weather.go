// To parse and unparse this JSON data, add this code to your project and do:
//
//    openWeather, err := UnmarshalOpenWeather(bytes)
//    bytes, err = openWeather.Marshal()

package remote

import "encoding/json"

func UnmarshalOpenWeather(data []byte) (OpenWeather, error) {
	var r OpenWeather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OpenWeather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OpenWeather struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int64   `json:"pressure"`
	Humidity int64   `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Sys struct {
	Type    int64   `json:"type"`
	ID      int64   `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
