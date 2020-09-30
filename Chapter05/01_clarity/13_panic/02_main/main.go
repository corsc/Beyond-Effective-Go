package main

import (
	"log"

	"github.com/corsc/Advanced-Go-Programming/Chapter05/01_clarity/13_panic/02_main/internal/config"
	"github.com/corsc/Advanced-Go-Programming/Chapter05/01_clarity/13_panic/02_main/internal/server"
)

func main() {
	cfg, err := config.Load("my-config.json")
	if err != nil {
		log.Fatalf("failed to load config with err: %s", err)
	}

	server := server.New(cfg)
	server.Start()
}
