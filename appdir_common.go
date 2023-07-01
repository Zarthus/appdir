package appdir

import (
	"errors"
	"os"
)

func tryEnvs(envs []string, fallback string) *string {
	var result string
	var found bool
	for _, env := range envs {
		result, found = os.LookupEnv(env)
		if found {
			return &result
		}
	}
	return &fallback
}

func handleEnvResult(value *string) (string, error) {
	if value == nil || *value == "" {
		return "", errors.New("could not find appropriate matching environment variables")
	}
	return *value, nil
}
