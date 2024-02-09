package viperConfig

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitViper(configFileName string) (*viper.Viper, error) {
    viperInstance := viper.New()
    viperInstance.SetConfigName(configFileName)
    viperInstance.SetConfigType("json") // Optionally, set config type if needed

    // Get the absolute path to the directory containing the configuration file
    configDir, err := filepath.Abs(filepath.Dir(".")) // Adjust if necessary
    if err != nil {
        return nil, err
    }

    // Set the directory to look for the config file
    viperInstance.AddConfigPath(configDir)

    // Find and read the config file
    err = viperInstance.ReadInConfig()
    if err != nil {
        return nil, fmt.Errorf("fatal error config file: %w", err)
    }

    return viperInstance, nil
}
