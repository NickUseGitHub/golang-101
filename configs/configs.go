package configs

import (
	"fmt"
	"os"
	"strings"
)

// Configs struct
type Configs struct {
	port       string
	dbHost     string
	dbPort     string
	dbName     string
	dbUsername string
	dbPassword string
	dbSslmode  string
}

// Initialize to init values from envs
func (c Configs) Initialize() Configs {
	c.port = "8080"
	if os.Getenv("PORT") != "" {
		c.port = os.Getenv("PORT")
	}

	c.dbHost = "localhost"
	if os.Getenv("DB_HOST") != "" {
		c.dbHost = os.Getenv("DB_HOST")
	}

	c.dbPort = "5432"
	if os.Getenv("DB_PORT") != "" {
		c.dbPort = os.Getenv("DB_PORT")
	}

	c.dbName = "todo"
	if os.Getenv("DB_NAME") != "" {
		c.dbName = os.Getenv("DB_NAME")
	}

	c.dbUsername = "postgres"
	if os.Getenv("DB_USERNAME") != "" {
		c.dbUsername = os.Getenv("DB_USERNAME")
	}

	c.dbPassword = "postgres"
	if os.Getenv("DB_PASSWORD") != "" {
		c.dbPassword = os.Getenv("DB_PASSWORD")
	}

	c.dbSslmode = "disable"
	if os.Getenv("DB_SSL_MODE") != "" {
		c.dbSslmode = os.Getenv("DB_SSL_MODE")
	}

	return c
}

// GetPort from config
func (c Configs) GetPort() string {
	return c.port
}

// GetDbConfigConnection get single string
func (c Configs) GetDbConfigConnection(configParam Configs) string {
	sliceEnvs := map[string]string{
		"host":     configParam.dbHost,
		"port":     configParam.dbPort,
		"dbname":   configParam.dbName,
		"user":     configParam.dbUsername,
		"password": configParam.dbPassword,
		"sslmode":  configParam.dbSslmode,
	}

	singleStr := func() string {
		strList := []string{}
		for key, value := range sliceEnvs {
			strList = append(strList, fmt.Sprintf("%s=%s", key, value))
		}

		return strings.Join(strList[:], " ")
	}()

	return singleStr
}
