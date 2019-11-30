// Package ido implements containers operation.
package ido

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Create creates an image directory.
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
	container, err := d.create(image)

	if err != nil {
		return "", err
	}

	tempFilePath := filepath.Join(tempDir, "temp.tar")
	if err = d.export(tempFilePath, container); err != nil {
		return "", err
	}

	if err = d.rm(container); err != nil {
		return "", err
	}

	if err = tarX(rootfsDir, tempFilePath); err != nil {
		return "", err
	}

	if err = os.Remove(tempFilePath); err != nil {
		return "", err
	}

	return tempDir, nil
}

// Run runs a container.
func Run(dir string, cmd string, volumes []string) error {
	// via https://ericchiang.github.io/post/containers-from-scratch/#creating-namespaces-with-unshare
	rootfsDir := filepath.Join(dir, "rootfs")

	for _, volume := range volumes {
		vDirs := strings.SplitN(volume, ":", 2)
		if len(vDirs) != 2 {
			return fmt.Errorf("inavlid volume: %s", volume)
		}

		hostDir, err := filepath.Abs(vDirs[0])
		if err != nil {
			return err
		}

		containerDir := filepath.Join(rootfsDir, vDirs[1])
		if err = os.MkdirAll(containerDir, 0750); err != nil {
			return err
		}

		mount := newShell("mount", "--bind", "-o", "ro", hostDir, containerDir)
		if _, err = mount.result(); err != nil {
			return err
		}

		umount := newShell("umount", containerDir)
		defer umount.run() // nolint: errcheck
	}

	err := unshareChroot(rootfsDir, cmd)
	if err != nil {
		return err
	}

	return nil
}

func mkRootfsDir(dir string) (rootfsDir string, err error) {
	rootfsDir = filepath.Join(dir, "rootfs")
	if err = os.Mkdir(rootfsDir, 0750); err != nil {
		return "", err
	}

	return rootfsDir, nil
}

func tarX(dir string, file string) (err error) {
	sh := newShell("tar", "-C", dir, "-xvf", file)
	if _, err = sh.result(); err != nil {
		return err
	}

	return nil
}

func unshareChroot(dir string, cmd string) (err error) {
	procDir := filepath.Join(dir, "proc")
	// via https://github.com/karelzak/util-linux/issues/648#issuecomment-404066455
	mount := newShell("mount", "--types", "proc", "proc", procDir)
	if _, err = mount.result(); err != nil {
		return err
	}

	umount := newShell("umount", procDir)
	defer umount.run() // nolint: errcheck

	uc := newShell("unshare", "--pid", "--fork", "--mount-proc="+procDir, "chroot", dir, cmd)
	if err = uc.run(); err != nil {
		return err
	}

	return nil
}
