package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JustinPras/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmd := command {
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatalf("error running command: %s", err)
	}
}

type state struct {
	cfg	*config.Config
}