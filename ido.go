package ido

import (
	"errors"
	"os/exec"
	"strings"
)

func Create(image string) (string, error) {
	b, err := exec.Command("docker", "create", image).CombinedOutput()
	output := strings.TrimSpace(string(b))
	if err != nil {
		return "", errors.New(output)
	}

	return output, nil
}
