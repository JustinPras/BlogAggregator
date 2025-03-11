package main

import (
	"fmt"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/JustinPras/BlogAggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), username); err != nil {
		return fmt.Errorf("that user doesn't exist: %w", err)
	}

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		return fmt.Errorf("that user already exists")
	}

	userParams := database.CreateUserParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		name,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("couldn't create new user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully!")
	fmt.Printf("User Data: %+v\n", user)
	return nil
}