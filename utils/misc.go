package utils

import (
	"os"
)

func ReadFromFile(target string) (string, error) {
	if _, err := os.Stat(target); err == nil { // check whether it is a path
		bytes, err := os.ReadFile(target)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	}

	return target, nil
}
