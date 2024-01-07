package siows

import (
	"fmt"

	"github.com/google/uuid"
)

// Config is a type that represents a configuration object.
// It contains the following fields:
// - id: a string representing the ID of the configuration
// - name: a string representing the name of the configuration
// - port: an unsigned 16-bit integer representing the port number
type Config struct {
	id   string
	name string
	port uint16
}

// Port returns the port number specified in the configuration.
func (c Config) Port() uint16 {
	return c.port
}

// Name returns the name of the configuration.
func (c Config) Name() string {
	return c.name
}

// ID returns the ID field of the Config object.
func (c Config) ID() string {
	return c.id
}

// WithID sets the ID field of the Config struct and returns the updated Config.
// The ID should be a string value.
func (c Config) WithID(id string) Config {
	c.id = id

	return c
}

// WithPort sets the port value of the Config struct and returns the modified Config.
func (c Config) WithPort(port uint16) Config {
	c.port = port

	return c
}

// WithName Set the name field of the config struct
func (c Config) WithName(name string) Config {
	c.name = name

	return c
}

// FmtString formats the Config struct as a string with specific format.
// It returns a formatted string containing the ID, Name, and Port of the Config.
// The format of the returned string is "ID: {id}, Name: {name}, Port: {port}".
// The values of id, name, and port are retrieved from the Config struct.
func (c Config) FmtString() string {
	return fmt.Sprintf("ID: %s, Name: %s, Port: %d", c.id, c.name, c.port)
}

// NewConfig creates a new instance of Config with default values.
func NewConfig() Config {
	return Config{
		id:   uuid.NewString(),
		port: 8080,
	}
}
