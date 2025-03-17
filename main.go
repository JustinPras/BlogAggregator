package main

import (
	"log"
	"os"
	"database/sql"
	"context"

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
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerDisplayFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))

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

func middlewareLoggedIn(handler func(s *state, cmd command, currentUser database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
        currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
        if err != nil {
            return err
        }

        return handler(s, cmd, currentUser)
    }
}