package main

import (
	"fmt"
	"github.com/JustinPras/BlogAggregator/internal/config"
)

func repl() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	cfg.SetUser("Justin")

	cfg, err = config.Read()
	if err != nil {
		return err
	}

	fmt.Println(cfg)
	return nil
}