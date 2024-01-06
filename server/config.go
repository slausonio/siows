package server

import "fmt"

type Config struct {
	id   string
	name string
	port uint16
}

func (c Config) Port() uint16 {
	return c.port
}

func (c Config) Name() string {
	return c.name
}

func (c Config) ID() string {
	return c.id
}

func (c Config) WithID(id string) Config {
	c.id = id

	return c
}

func (c Config) WithPort(port uint16) Config {
	c.port = port

	return c
}

func (c Config) WithName(name string) Config {
	c.name = name

	return c
}

func (c Config) FmtString() string {
	return fmt.Sprintf("ID: %s, Name: %s, Port: %d", c.id, c.name, c.port)
}

func NewConfig() Config {
	return Config{
		port: 8080,
	}
}
