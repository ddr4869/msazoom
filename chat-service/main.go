package main

import (
	"log"

	"github.com/ddr4869/msazoom/chat-service/config"
	"github.com/ddr4869/msazoom/chat-service/internal"
)

func main() {
	cfg := config.Init()
	router, err := internal.NewRestController(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}
	router.Start()
}
