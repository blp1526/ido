package ido

import (
	"os"
	"path/filepath"
)

func MkRootfs(dir string) error {
	rootfsDir := filepath.Join(dir, "rootfs")
	err := os.Mkdir(rootfsDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func TarX(dir string, file string) (string, error) {
	result, err := command("tar", "-C", dir, "-xvf", file)
	if err != nil {
		return "", err
	}

	return result, nil
}

func Chroot(dir string, cmd string) (string, error) {
	result, err := command("chroot", dir, cmd)
	if err != nil {
		return "", err
	}

	return result, nil
}
