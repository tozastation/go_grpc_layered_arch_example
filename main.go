package main

import (
	"github.com/tozastation/gRPC-Training-Golang/interfaces/di"
	rpc_ping "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/ping"
	rpc_user "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/user"
	rpc_weather "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/weather"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalln(err)
	}

	// Dependency Injection
	server := di.InitializeServer()
	weather := di.InitializeWeather()
	ping := di.InitializePing()
	user := di.InitializeUser()
	// register RPC
	rpc_user.RegisterUsersServer(server, user)
	rpc_ping.RegisterCheckServer(server, ping)
	rpc_weather.RegisterWeathersServer(server, weather)
	server.Serve(listenPort)
}
