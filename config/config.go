package config

import (
	"context"
	"os"
	"strconv"
	"time"
)

// Web defines the web configuration
type Web struct {
	APIHost            string
	APIReadTimeout     time.Duration
	APIWriteTimeout    time.Duration
	APIShutDownTimeout time.Duration
	DebugHost          string
	DebugMode          bool
}

// Config defines the configuration
type Config struct {
	Web Web
}

// New generates a new Configuration type
func New(ctx context.Context) *Config {
	cfg := &Config{
		Web: Web{
			APIHost:            getEnv("TURF_API_HOST", "0.0.0.0:3000"),
			APIReadTimeout:     getEnvAsTimeDuration("TURF_API_READ_TIMEOUT", "5s"),
			APIWriteTimeout:    getEnvAsTimeDuration("TURF_API_WRITE_TIMEOUT", "5s"),
			APIShutDownTimeout: getEnvAsTimeDuration("TURF_API_SHUTDOWN_TIMEOUT", "5s"),
			DebugHost:          getEnv("TURF_API_DEBUG_HOST", "0.0.0.0:4000"),
			DebugMode:          getEnvAsBool("DEBUG_MODE", true),
		},
	}

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsTimeDuration(name string, defaultVal string) time.Duration {
	def, _ := time.ParseDuration(defaultVal)
	if value, ok := os.LookupEnv(name); ok {
		d, err := time.ParseDuration(value)
		if err != nil {
			return def
		}
		return d
	}

	return def
}
