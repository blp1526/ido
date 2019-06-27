package ido

type docker struct{}

func newDocker() *docker {
	return &docker{}
}

func (d docker) create(image string) (container string, err error) {
	container, err = command("docker", "create", image)
	if err != nil {
		return "", err
	}

	return container, nil
}

func (d docker) export(outputDir string, container string) (err error) {
	_, err = command("docker", "export", "-o", outputDir, container)
	if err != nil {
		return err
	}

	return nil
}
