package main

import (
	"errors"
)

type command struct {
    Name string
    Args []string
}
type commands struct {
    registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
    // store f in the map using name as the key
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
    // look up cmd.Name in the map
    // if not found, return an error
    // if found, call it with s and cmd
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	
	return f(s, cmd)
}