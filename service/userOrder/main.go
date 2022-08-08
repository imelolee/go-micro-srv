package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"userOrder/handler"
	"userOrder/model"
	pb "userOrder/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "userorder"
	version = "latest"
)

func main() {
	model.InitRedis()
	model.InitDb()

	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register handler
	err := pb.RegisterUserOrderHandler(srv.Server(), new(handler.UserOrder))
	if err != nil {
		fmt.Println("RegisterUserOrderHandler err: ", err)
		return
	}
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
