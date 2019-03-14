package main

import (
	"fmt"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/di"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/handler"
	rpc_ping "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/ping"
	rpc_user "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/user"
	rpc_weather "github.com/tozastation/go-grpc-ddd-example/interfaces/rpc/weather"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}
	// Establish DB Connection
	// handler.DB, err = handler.OpenDBConnection()
	// if err != nil {
	// 	panic(err)
	// }
	// Dependency Injection
	server := di.InitializeServer()
	weather := di.InitializeWeather()
	ping := di.InitializePing()
	//user := di.InitializeUser()
	// register RPC
	//rpc_user.RegisterUsersServer(server, user)
	rpc_ping.RegisterCheckServer(server, ping)
	rpc_weather.RegisterWeathersServer(server, weather)
	fmt.Println("Listen on 3001")
	server.Serve(listenPort)
}
