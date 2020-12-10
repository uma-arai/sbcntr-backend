package utils

import (
	"os"
)

// APIConfig ...
type APIConfig struct {
	HeaderValue struct {
		ClientID string
	}
}

// ConfigDB ...
type ConfigDB struct {
	MySQL struct {
		DBMS     string
		Protocol string
		Username string
		Password string
		DBName   string
	}
}

// NewAPIConfig ...
func NewAPIConfig() *APIConfig {
	config := new(APIConfig)
	config.HeaderValue.ClientID = os.Getenv("CNAPP_CLIENT_ID_HEADER")

	return config
}

// NewConfigDB ...
func NewConfigDB() *ConfigDB {
	config := new(ConfigDB)

	config.MySQL.DBMS = "mysql"
	config.MySQL.Protocol = "tcp(" + os.Getenv("CNAPP_DB_HOST") + ":3306)"
	config.MySQL.Username = os.Getenv("CNAPP_DB_USERNAME")
	config.MySQL.Password = os.Getenv("CNAPP_DB_PASSWORD")
	config.MySQL.DBName = os.Getenv("CNAPP_DB_NAME")

	return config
}
