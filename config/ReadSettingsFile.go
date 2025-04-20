package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadSettingFile() (Settings, error) {
	var settings Settings

	path, err := getConfigPath()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	yamlFile, err := os.ReadFile(path)

	if err != nil {
		return settings, err
	}

	if err = yaml.Unmarshal(yamlFile, &settings); err != nil {
		return settings, err
	}

	return settings, nil
}
