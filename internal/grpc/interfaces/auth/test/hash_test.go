package test

import (
	"fmt"
	srv "github.com/tozastation/go-grpc-ddd-example/domain/service"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/auth"
	rpc_user "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/user"
	"testing"
)

func TestHashed(t *testing.T) {
	before := &rpc_user.PostUser{
		UserID:   "tozastation",
		CityName: "Hakodate",
		Password: "Test@1234",
		Name:     "tozastation",
	}
	fmt.Println("Base Password: " + before.Password)
	expectedPassword, _ := auth.Hashed("Test@1234")
	actual, _ := srv.PostUserToDB(before)
	if err := auth.CheckHash(actual.Password, "Test@1234"); err != nil {
		t.Errorf(err.Error())
	}
	if err := auth.CheckHash(expectedPassword, "Test@1234"); err != nil {
		t.Errorf(err.Error())
	}
}
