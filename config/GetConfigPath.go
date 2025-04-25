package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
	localPath := "config.yml"
	if _, err := os.Stat(localPath); err == nil {
		return localPath, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(home, ".config", "commitwise", "config.yml")

	if _, err := os.Stat(configPath); err == nil {
		return configPath, nil
	}

	return "", fmt.Errorf("settings file not found in either %s or %s", localPath, configPath)
}
