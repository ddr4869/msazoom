package main

import (
	"log"

	"github.com/ddr4869/msazoom/user-service/config"
	"github.com/ddr4869/msazoom/user-service/grpc"
	"github.com/ddr4869/msazoom/user-service/internal"
)

func main() {
	cfg := config.Init()
	router, err := internal.NewRestController(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}
	// setup grpc client
	conn := grpc.NewMessageClient()
	defer conn.Close()
	router.Start()
}
