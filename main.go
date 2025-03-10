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

	s := state{
		cfgPtr: &cfg,
	}

	//create map of handler functions
	handlerFunctions := make(map[string]func(*state, command) error)

	commands := commands{
		handlerFunctions: handlerFunctions,
	}

	commands.register("login", handlerLogin)
	

	if len(os.Args) < 2 {
		log.Fatalf("no arguments provided")
	}

	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = commands.run(&s, cmd)
	if err != nil {
		log.Fatalf("error running command: %s", err)
	}
	
}

type state struct {
	cfgPtr	*config.Config
}

type command struct {
	name	string
	args	[]string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	username := cmd.args[0]

	s.cfgPtr.SetUser(username)
	fmt.Println("User has been set")
	return nil
}

type commands struct {
	handlerFunctions map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlerFunctions[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.handlerFunctions[cmd.name]
	if !ok {
		return fmt.Errorf("Command does not exist")
	}

	err := f(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

