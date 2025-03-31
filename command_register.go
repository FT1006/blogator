package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/FT1006/blogator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("no user name")
	}

	userName := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), userName); err == nil {
		fmt.Println("Error: User already exists")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	err = s.config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Printf("User %s created successfully\n", newUser.Name)
	fmt.Printf("ID: %s\n", newUser.ID)
	fmt.Printf("Created at: %s\n", newUser.CreatedAt)
	fmt.Printf("Updated at: %s\n", newUser.UpdatedAt)

	return nil
}
