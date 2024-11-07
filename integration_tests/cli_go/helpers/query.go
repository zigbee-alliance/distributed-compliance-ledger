package helpers

func Query(module, command string, queryArgs ...string) (string, error) {
	args := []string{"query", module, command}
	args = append(args, queryArgs...)

	output, err := Command(args...)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
