// Package config provides environment-based configuration management.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Base holds common configuration shared across modules.
type Base struct {
	ServiceName string
	Environment string
	LogLevel    string
	HTTPPort    int
	ShutdownTimeout time.Duration
}

// LoadBase reads standard environment variables with sensible defaults.
func LoadBase() Base {
	return Base{
		ServiceName:     getEnv("SERVICE_NAME", "go-mastery-roadmap"),
		Environment:     getEnv("ENVIRONMENT", "development"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		HTTPPort:        getEnvInt("HTTP_PORT", 8080),
		ShutdownTimeout: getEnvDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}

// Validate ensures required fields are set for production.
func (b Base) Validate() error {
	if b.ServiceName == "" {
		return fmt.Errorf("SERVICE_NAME is required")
	}
	return nil
}
