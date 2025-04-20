package utils

import (
	"os"
	"os/exec"
)

func Commit(message string) error {
	file, err := os.CreateTemp("", "commitwise_commitmsg_tmp")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(message)); err != nil {
		return err
	}

	cmd := exec.Command("git", "commit", "--edit", "-F", file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	return err
}
