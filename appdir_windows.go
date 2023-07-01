//go:build windows

package appdir

func GetApplicationDirectory() (string, error) {
	return handleEnvResult(tryEnvs([]string{"ProgramFiles", "HOME"}, ""))
}

func GetUserConfigDirectory() (string, error) {
	return handleEnvResult(tryEnvs([]string{"APPDATA", "HOME"}, ""))
}
