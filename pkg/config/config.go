package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
    ExportPath string
}

var Config AppConfig

func InitAppConfig() {
    fmt.Println("Initializing config...")
    
    value, exists := os.LookupEnv("EXPORT_PATH")
    fmt.Printf("EXPORT_PATH exists: %v, value: %s\n", exists, value)
    
    Config = AppConfig{
        ExportPath: getEnvAsString("EXPORT_PATH", "./data"),
    }
    fmt.Println("Config loaded:", Config)
}

func getEnvAsString(key string, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}