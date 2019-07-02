package ido

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Create(image string) (tempDir string, err error) {
	// via https://github.com/opencontainers/runc/blob/6cccc1760d57d9e1bc856b96eeb7ee02b7b8101d/README.md#using-runc
	tempDir, err = ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}

	rootfsDir, err := mkRootfsDir(tempDir)
	if err != nil {
		return "", err
	}

	d := newDocker()
	//
	container, err := d.create(image)
	if err != nil {
		return "", err
	}

	tempFilePath := filepath.Join(tempDir, "temp.tar")
	err = d.export(tempFilePath, container)
	if err != nil {
		return "", err
	}
	err = d.rm(container)
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

func Run(dir string, cmd string) error {
	// via https://ericchiang.github.io/post/containers-from-scratch/#creating-namespaces-with-unshare
	rootfsDir := filepath.Join(dir, "rootfs")
	err := unshareChroot(rootfsDir, cmd)
	if err != nil {
		return err
	}

	err = os.RemoveAll(dir)
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
	procDir := filepath.Join(dir, "proc")
	// via https://github.com/karelzak/util-linux/issues/648#issuecomment-404066455
	mount := newShell("mount", "-t", "proc", "proc", procDir)
	err = mount.run()
	if err != nil {
		return err
	}

	sh := newShell("unshare", "--pid", "--fork", "--mount-proc="+procDir, "chroot", dir, cmd)
	sh.cmd.Stdout = os.Stdout
	sh.cmd.Stdin = os.Stdin
	sh.cmd.Stderr = os.Stderr

	err = sh.run()
	if err != nil {
		return err
	}

	umount := newShell("umount", procDir)
	err = umount.run()
	if err != nil {
		return err
	}

	return nil
}
