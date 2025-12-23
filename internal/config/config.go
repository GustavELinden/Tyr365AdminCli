// Package config provides a singleton configuration manager using Viper.
// Initialize once at startup, then use GetConfig() throughout the application.
package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	instance *viper.Viper
	once     sync.Once
	initErr  error
)

// DefaultConfigPaths contains the default paths to search for config files
var DefaultConfigPaths = []string{
	"/root/condigurationFolder/",
	"~/condigurationFolder/",
	".",
}

// Initialize sets up the config singleton. Call this once at application startup.
// If called multiple times, subsequent calls are no-ops.
func Initialize(configFileName string) error {
	once.Do(func() {
		instance = viper.New()
		instance.SetConfigName(configFileName)
		instance.SetConfigType("json")

		for _, path := range DefaultConfigPaths {
			instance.AddConfigPath(path)
		}

		if err := instance.ReadInConfig(); err != nil {
			initErr = fmt.Errorf("failed to read config file: %w", err)
			return
		}
	})

	return initErr
}

// Get returns the config singleton instance.
// Panics if Initialize() was not called first or failed.
func Get() *viper.Viper {
	if instance == nil {
		panic("config not initialized: call config.Initialize() first")
	}
	return instance
}

// MustGet returns the config singleton, initializing with defaults if needed.
// This is useful for packages that may be used before main() initializes config.
func MustGet() *viper.Viper {
	if instance == nil {
		if err := Initialize("config.json"); err != nil {
			panic(fmt.Sprintf("failed to initialize config: %v", err))
		}
	}
	return instance
}

// GetString is a convenience method to get a string value from config
func GetString(key string) string {
	return Get().GetString(key)
}

// GetInt is a convenience method to get an int value from config
func GetInt(key string) int {
	return Get().GetInt(key)
}

// GetBool is a convenience method to get a bool value from config
func GetBool(key string) bool {
	return Get().GetBool(key)
}

// IsSet checks if a key is set in the config
func IsSet(key string) bool {
	return Get().IsSet(key)
}
