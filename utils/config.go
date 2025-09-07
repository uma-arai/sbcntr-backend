package utils

import (
	"os"
	"strings"
)

// APIConfig ...
type APIConfig struct {
	Env         string // "development", "production"
	HeaderValue struct {
		ClientID string
	}
	EnableTracing bool
}

// ConfigDB ...
type ConfigDB struct {
	Postgres struct {
		DBMS     string
		Username string
		Password string
		DBName   string
	}
}

// NewAPIConfig ...
func NewAPIConfig() *APIConfig {
	config := new(APIConfig)
	config.HeaderValue.ClientID = os.Getenv("SBCNTR_CLIENT_ID_HEADER")

	// 環境変数[SBCNTR_ENABLE_TRACING]を見てトレースを有効にする。対応しているTracingはAWS_XRAYのみ。
	enableKey := os.Getenv("SBCNTR_ENABLE_TRACING")
	if strings.ToLower(enableKey) == "true" || enableKey == "1" {
		config.EnableTracing = true
	} else {
		config.EnableTracing = false
	}

	config.Env = os.Getenv("APP_ENV")
	if config.Env == "" {
		config.Env = "development"
	}

	return config
}

// NewConfigDB ...
func NewConfigDB() *ConfigDB {
	config := new(ConfigDB)

	config.Postgres.DBMS = "postgres"
	config.Postgres.Username = os.Getenv("DB_USERNAME")
	config.Postgres.Password = os.Getenv("DB_PASSWORD")
	config.Postgres.DBName = os.Getenv("DB_NAME")

	return config
}
