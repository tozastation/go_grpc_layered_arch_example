// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tozastation/gRPC-Training-Golang/domain/service (interfaces: IWeatherService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	weather "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/weather"
	reflect "reflect"
)

// MockIWeatherService is a mock of IWeatherService interface
type MockIWeatherService struct {
	ctrl     *gomock.Controller
	recorder *MockIWeatherServiceMockRecorder
}

// MockIWeatherServiceMockRecorder is the mock recorder for MockIWeatherService
type MockIWeatherServiceMockRecorder struct {
	mock *MockIWeatherService
}

// NewMockIWeatherService creates a new mock instance
func NewMockIWeatherService(ctrl *gomock.Controller) *MockIWeatherService {
	mock := &MockIWeatherService{ctrl: ctrl}
	mock.recorder = &MockIWeatherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIWeatherService) EXPECT() *MockIWeatherServiceMockRecorder {
	return m.recorder
}

// GetWeather mocks base method
func (m *MockIWeatherService) GetWeather(arg0 context.Context, arg1 string) (*weather.Weather, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeather", arg0, arg1)
	ret0, _ := ret[0].(*weather.Weather)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeather indicates an expected call of GetWeather
func (mr *MockIWeatherServiceMockRecorder) GetWeather(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeather", reflect.TypeOf((*MockIWeatherService)(nil).GetWeather), arg0, arg1)
}
