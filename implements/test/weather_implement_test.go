package test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	mock "github.com/tozastation/go-grpc-ddd-example/implements/mock"
	rpc_weather "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/weather"
	"testing"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockImp := mock.NewMockIWeatherImplement(ctrl)
	req := rpc_weather.GetRequest{
		CityName: "Hakodate",
	}
	res := rpc_weather.GetResponse{
		Weather: &rpc_weather.Weather{
			ID:          0,
			CityName:    "Hakodate",
			TempMax:     100,
			TempMin:     0,
			Wind:        100,
			Type:        "台風",
			Description: "強風注意",
		},
	}
	mockImp.EXPECT().Get(
		gomock.Any(),
		&req,
	).Return(&res, nil)
	testGet(t, mockImp)
}

func testGet(t *testing.T, server rpc_weather.WeathersServer) {
	t.Helper()
	req := rpc_weather.GetRequest{
		CityName: "Hakodate",
	}
	res, _ := server.Get(context.Background(), &req)
	expected := res.GetWeather()
	actual := &rpc_weather.Weather{
		ID:          0,
		CityName:    "Hakodate",
		TempMax:     100,
		TempMin:     0,
		Wind:        100,
		Type:        "台風",
		Description: "強風注意",
	}
	if diff := cmp.Diff(*actual, *expected); diff != "" {
		t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
	}
	t.Log("Reply : ", res.Weather)
}
