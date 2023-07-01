//go:build darwin

package appdir

import (
	"errors"
	"fmt"
	"os"
)

func GetApplicationDirectory() (string, error) {
	home, found := os.LookupEnv("HOME")

	if !found {
		return "", errors.New("Could not find home directory")
	}

	return fmt.Sprintf("%s/Library/Application Support", home), nil
}

func GetUserConfigDirectory() (string, error) {
	home, found := os.LookupEnv("HOME")

	if !found {
		return "", errors.New("Could not find home directory")
	}

	return fmt.Sprintf("%s/Library/Preferences", home), nil
}
