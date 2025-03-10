package main

import (
	"fmt"
	"github.com/JustinPras/BlogAggregator/internal/config"
)

func justin() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	// fmt.Printf("Printing Config:\n%s\n", cfg)

	cfg.SetUser("Justin")

	cfg, err = config.Read()
	if err != nil {
		return err
	}

	fmt.Printf("Printing Config:\n%s\n", cfg)

	return nil
}