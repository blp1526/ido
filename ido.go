package ido

import (
	"os"
	"path/filepath"
)

func Run() error {
	wd := filepath.Join("/tmp", "tmp")
	rootfsDir, err := MkRootfs(wd)
	if err != nil {
		return err
	}

	result, err := Create("busybox")
	if err != nil {
		return err
	}

	result, err = Export(result, wd)
	if err != nil {
		return err
	}

	_, err = TarX(rootfsDir, result)
	if err != nil {
		return err
	}

	_, err = Chroot(rootfsDir, "/bin/sh")
	if err != nil {
		return err
	}

	return nil
}

func MkRootfs(dir string) (string, error) {
	rootfsDir := filepath.Join(dir, "rootfs")
	err := os.Mkdir(rootfsDir, 0755)
	if err != nil {
		return "", err
	}

	return rootfsDir, nil
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
