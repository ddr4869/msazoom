package main

import (
	"log"

	"github.com/ddr4869/msazoom/message-service/config"
	"github.com/ddr4869/msazoom/message-service/grpc"
	"github.com/ddr4869/msazoom/message-service/internal"
)

func main() {
	cfg := config.Init()
	router, err := internal.NewRestController(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}

	go grpc.SetGrpcServer(cfg)
	router.Start()
}
