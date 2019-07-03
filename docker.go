package ido

type docker struct{}

func newDocker() *docker {
	return &docker{}
}

func (d docker) create(image string) (container string, err error) {
	sh := newShell("docker", "create", image)
	container, err = sh.result()
	if err != nil {
		return "", err
	}

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
