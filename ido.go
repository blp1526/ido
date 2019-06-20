package ido

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func MkRootfs(dir string) error {
	rootfsDir := filepath.Join(dir, "rootfs")
	err := os.Mkdir(rootfsDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func Create(image string) (string, error) {
	result, err := command("docker", "create", image)
	if err != nil {
		return "", err
	}

	return result, nil
}

func Export(container string, output string) (string, error) {
	result, err := command("docker", "export", "-o", output, container)
	if err != nil {
		return "", err
	}

	return result, nil
}

func TarX(dir string, file string) (string, error) {
	result, err := command("tar", "-C", dir, "-xvf", file)
	if err != nil {
		return "", err
	}

	return result, nil
}
