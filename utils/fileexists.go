package utils

import (
	"errors"
	"os"
)

// utility function if a file exists or not
func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
