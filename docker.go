package ido

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
