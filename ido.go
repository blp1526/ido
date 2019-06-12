package ido

import (
	"fmt"
	"os/exec"
	"strings"
)

func command(name string, arg ...string) (result string, err error) {
	b, err := exec.Command(name, arg...).CombinedOutput()
	result = strings.TrimSpace(string(b))
	if err != nil {
		words := append([]string{name}, arg[0:]...)
		return "", fmt.Errorf("%s: %s", strings.Join(words, " "), result)
	}

	return result, nil
}

func Create(image string) (string, error) {
	result, err := command("docker", "create", image)
	if err != nil {
		return "", err
	}

	return result, nil
}
