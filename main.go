package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FT1006/blogator/internal/config"
)

type state struct {
	Config *config.Config
}

func newState(cfg *config.Config) *state {
	return &state{
		Config: cfg,
	}
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)	
	}

	currentState := newState(&cfg)

	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1)  // Exit with non-zero status code to indicate an error
	}

	var cmdMap commands
	cmdMap.CommandMap = make(map[string]func(*state, command) error)
	cmdMap.register("login", handlerLogin)

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmdMap.run(currentState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
