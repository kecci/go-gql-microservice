package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/kecci/go-gql-microservice/internal/config"
)

const repo = "go-gql-microservice"

func main() {

	log.Println("Initialize configuration")
	cfg, err := config.New(repo)
	if err != nil {
		log.Fatalf("failed to init the config: %v", err)
	}

	startApp(cfg)
}
