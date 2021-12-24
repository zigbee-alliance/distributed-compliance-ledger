package cli

import (
	"io/ioutil"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ReadFromFile(target string) (string, error) {
	if _, err := os.Stat(target); err == nil { // check whether it is a path
		bytes, err := ioutil.ReadFile(target)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	} else { // else return as is
		return target, nil
	}
}

func HandleError(err error) error {
	if err == nil {
		return nil
	}
	s, ok := status.FromError(err)
	if !ok {
		return err
	}
	if s.Code() != codes.NotFound {
		return err
	}
	return nil
}
