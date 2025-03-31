package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("no user name")
	}

	userName := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		fmt.Printf("You can't login to an account that doesn't exist!: %v\n", err)
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	err = s.config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println(s.config.CurrentUserName)
	return nil
}
