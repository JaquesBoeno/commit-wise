package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigPath() (string, error) {
	devPath := "config.yml"
	if _, err := os.Stat(devPath); err == nil {
		return devPath, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	prodPath := filepath.Join(home, ".config", "commitwise", "config.yml")

	if _, err := os.Stat(prodPath); err == nil {
		return prodPath, nil
	}

	return "", fmt.Errorf("settings file not found in either %s or %s", devPath, prodPath)
}
