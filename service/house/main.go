package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"house/handler"
	"house/model"
	pb "house/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "house"
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
	err := pb.RegisterHouseHandler(srv.Server(), new(handler.House))
	if err != nil {
		fmt.Println("RegisterUserHandler err: ", err)
		return
	}
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
