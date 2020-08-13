package configs

import "os"

// Configs struct
type Configs struct {
	port string
}

// Initialize to init values from envs
func (c Configs) Initialize() Configs {
	os.Setenv("PORT", "8080")
	c.port = os.Getenv("PORT")

	return c
}

// GetPort from config
func (c Configs) GetPort() string {
	return c.port
}
