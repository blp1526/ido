package ido

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Create(image string) (tempDir string, err error) {
	tempDir, err = ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}

	rootfsDir, err := mkRootfsDir(tempDir)
	if err != nil {
		return "", err
	}

	d := newDocker()
	container, err := d.create(image)
	if err != nil {
		return "", err
	}

	tempFilePath := filepath.Join(tempDir, "temp.tar")
	err = d.export(tempFilePath, container)
	if err != nil {
		return "", err
	}

	err = tarX(rootfsDir, tempFilePath)
	if err != nil {
		return "", err
	}

	err = os.Remove(tempFilePath)
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

func Run(cmd string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	rootfsDir := filepath.Join(wd, "rootfs")
	err = unshareChroot(rootfsDir, cmd)
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
	sh := newShell("tar", "-C", dir, "-xvf", file)
	err = sh.run()
	if err != nil {
		return err
	}

	return nil
}

func unshareChroot(dir string, cmd string) (err error) {
	sh := newShell("unshare", "--pid", "--fork", "chroot", dir, cmd)
	sh.cmd.Stdout = os.Stdout
	sh.cmd.Stdin = os.Stdin
	sh.cmd.Stderr = os.Stderr

	err = sh.run()
	if err != nil {
		return err
	}

	return nil
}
