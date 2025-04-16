package utils

import (
	"os"
	"strconv"
)

type DBConfig struct {
	EnableLogging bool
	MaxIdleConns  int
	MaxOpenConns  int
}

var SQLiteConfig DBConfig

func InitDBConfig() DBConfig {
	return DBConfig{
		EnableLogging: getEnvAsBool("ENABLE_LOGGING", false),
		MaxIdleConns:  getEnvAsInt("MAX_IDLE_CONNS", 0),
		MaxOpenConns:  getEnvAsInt("MAX_OPEN_CONNS", 0),
	}
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		switch value {
		case "true", "1", "yes":
			return true
		case "false", "0", "no":
			return false
		}
	}
	return fallback
}
