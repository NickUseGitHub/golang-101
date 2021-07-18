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
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "todo")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_SSL_MODE", "disable")

	c.port = os.Getenv("PORT")
	c.dbHost = os.Getenv("DB_HOST")
	c.dbPort = os.Getenv("DB_PORT")
	c.dbName = os.Getenv("DB_NAME")
	c.dbUsername = os.Getenv("DB_USERNAME")
	c.dbPassword = os.Getenv("DB_PASSWORD")
	c.dbSslmode = os.Getenv("DB_SSL_MODE")

	fmt.Println(c.dbHost)

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
