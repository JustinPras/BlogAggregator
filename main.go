package main

import (
	"log"
	"os"
	"database/sql"

	"github.com/JustinPras/BlogAggregator/internal/config"
	"github.com/JustinPras/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db	*database.Queries
	cfg	*config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	dbQueries := database.New(db)

	programState := state{
		db:		dbQueries,
		cfg: 	&cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAggregator)
	cmds.register("addfeed", handlerAddFeed)

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

