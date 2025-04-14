package utils

import (
	"os"
	"os/exec"
)

func Commit(message string) (string, error) {
	file, err := os.CreateTemp("", "commitwise_commitmsg_tmp")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(message)); err != nil {
		return "", err
	}

	cmd := exec.Command("git", "commit", "-F")
	cmd.Args = append(cmd.Args, file.Name())

	result, err := cmd.CombinedOutput()

	return string(result), err
}
