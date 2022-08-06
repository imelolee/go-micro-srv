package main

import (
	"house/handler"
	pb "house/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "house"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterHouseHandler(srv.Server(), new(handler.House))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
