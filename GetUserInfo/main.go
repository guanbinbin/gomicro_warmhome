package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gomicro/GetUserInfo/handler"

	"github.com/micro/go-grpc"
	example "gomicro/GetUserInfo/proto/example"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
