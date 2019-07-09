package ido

import "strings"

type docker struct{}

func newDocker() *docker {
	return &docker{}
}

func (d docker) create(image string) (container string, err error) {
	sh := newShell("docker", "create", image)
	result, err := sh.result()
	if err != nil {
		return "", err
	}

	// --- result example ---
	// Unable to find image 'foo:latest' locally
	// latest: Pulling from library/foo
	// abcdefghijkl: Pulling fs layer
	// abcdefghijkl: Download complete
	// abcdefghijkl: Pull complete
	// Digest: sha256:abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqrstuvwxyz01
	// Status: Downloaded newer image for foo:latest
	// abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqrstuvwxyz01

	lines := strings.Split(result, "\n")
	container = lines[len(lines)-1]
	return container, nil
}

func (d docker) export(outputDir string, container string) (err error) {
	sh := newShell("docker", "export", "-o", outputDir, container)
	_, err = sh.result()
	if err != nil {
		return err
	}

	return nil
}

func (d docker) rm(container string) (err error) {
	sh := newShell("docker", "rm", container)
	_, err = sh.result()
	if err != nil {
		return err
	}

	return nil
}
