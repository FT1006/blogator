package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	CommandMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.CommandMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.CommandMap[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return f(s, cmd)
}