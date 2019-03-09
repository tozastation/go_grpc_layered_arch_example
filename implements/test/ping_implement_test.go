package test

import (
	"context"
	"github.com/golang/mock/gomock"
	mock "github.com/tozastation/gRPC-Training-Golang/implements/mock"
	rpc_ping "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/ping"
	"testing"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockImp := mock.NewMockIPingImplement(ctrl)
	req := rpc_ping.Empty{}
	res := rpc_ping.Pong{
		Reply: "Pong",
	}
	mockImp.EXPECT().Ping(
		gomock.Any(),
		&req,
	).Return(&res, nil)
	testPing(t, mockImp)
}

func testPing(t *testing.T, server rpc_ping.CheckServer) {
	t.Helper()
	req := rpc_ping.Empty{}
	res, _ := server.Ping(context.Background(), &req)
	expected := res.Reply
	actual := "Pong"
	if expected != actual {
		t.Errorf("Failed")
	}
	t.Log("Reply : ", res.Reply)
}
