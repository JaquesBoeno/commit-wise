package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type question struct {
	Id        string     `yaml:"id"`
	Type      string     `yaml:"type"`
	Label     string     `yaml:"label"`
	Min       int        `yaml:"min"`
	Max       int        `yaml:"max"`
	Options   []option   `yaml:"options"`
	Questions []question `yaml:"questions"`
}
type option struct {
	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
}

type Settings struct {
	TemplateCommit string     `yaml:"TemplateCommit"`
	Questions      []question `yaml:"Questions"`
}

func getConfigPath() (string, error) {
	devPath := "config.yml"
	if _, err := os.Stat(devPath); err == nil {
		return devPath, nil
	}

	// Caminho global (modo prod)
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	prodPath := filepath.Join(home, ".config", "commitwise", "config.yml")

	if _, err := os.Stat(prodPath); err == nil {
		return prodPath, nil
	}

	return "", fmt.Errorf("arquivo de configuração não encontrado em %s nem em %s", devPath, prodPath)
}

func ReadSettingFile() (Settings, error) {
	var settings Settings

	path, err := getConfigPath()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return settings, err
	}

	if err = yaml.Unmarshal(data, &settings); err != nil {
		return settings, err
	}

	return settings, nil
}
