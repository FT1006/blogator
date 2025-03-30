package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no user name")
	}

	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println(s.Config.CurrentUserName)
	return nil
}
