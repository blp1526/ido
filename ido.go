package ido

import (
	"os"
	"path/filepath"
	"strconv"
)

func Create(image string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	rootfsDir, err := mkRootfsDir(wd)
	if err != nil {
		return err
	}

	d := newDocker()

	container, err := d.create(image)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(wd, "tmp.tar")
	err = d.export(outputDir, container)
	if err != nil {
		return err
	}

	err = tarX(rootfsDir, outputDir)
	if err != nil {
		return err
	}

	err = os.Remove(outputDir)
	if err != nil {
		return err
	}

	return nil
}

func Run(cmd string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	rootfsDir := filepath.Join(wd, "rootfs")
	err = chroot(rootfsDir, "/bin/sh")
	if err != nil {
		return err
	}

	return nil
}

func Attach(pid string) error {
	_, err := strconv.Atoi(pid)
	if err != nil {
		return err
	}

	return nil
}

func mkRootfsDir(dir string) (rootfsDir string, err error) {
	rootfsDir = filepath.Join(dir, "rootfs")
	err = os.Mkdir(rootfsDir, 0755)
	if err != nil {
		return "", err
	}

	return rootfsDir, nil
}

func tarX(dir string, file string) (err error) {
	_, err = command("tar", "-C", dir, "-xvf", file)
	if err != nil {
		return err
	}

	return nil
}

func chroot(dir string, cmd string) (err error) {
	_, err = command("chroot", dir, cmd)
	if err != nil {
		return err
	}

	return nil
}
