package utils

import (
	"os/exec"
)

func RunCommand(args ...string) (bool, error) {
	cmd := exec.Command(args[0], args[1:]...)
    _, err := cmd.Output()

    if err != nil {
		return false, err
    }
	return true, nil
}