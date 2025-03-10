package main

import (
	"fmt"
	"log"

	"github.com/JustinPras/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Printing config: %+v\n", cfg)

	err = cfg.SetUser("Justin")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Printing config again: %+v\n", cfg)
}