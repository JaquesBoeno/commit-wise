package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadSettingFile(path string) (Settings, error) {
	var settings Settings

	yamlFile, err := os.ReadFile(path)

	if err != nil {
		return settings, err
	}

	if err = yaml.Unmarshal(yamlFile, &settings); err != nil {
		return settings, err
	}

	return settings, nil
}
