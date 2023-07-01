//go:build linux

package appdir

import "errors"

func GetApplicationDirectory() (string, error) {
	return "", errors.New("Not yet implemented for linux")
}

func GetUserConfigDirectory() (string, error) {
	return handleEnvResult(tryEnvs([]string{"XDG_CONFIG_HOME", "HOME"}, ""))
}
