package cli

import (
	"os/exec"

	"github.com/cosmos/cosmos-sdk/types/errors"
)

func Execute(args ...string) (string, error) {
	cmd := exec.Command(CliBinaryName, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(out))
	}

	return string(out), err
}
