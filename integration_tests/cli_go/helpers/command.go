package helpers

import (
	"os/exec"
)

const CliBinaryName = "dcld"

func Command(args ...string) ([]byte, error) {
	cmd := exec.Command(CliBinaryName, args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return out, err
}
