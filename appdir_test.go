package appdir

import (
	"os"
	"regexp"
	"runtime"
	"testing"
)

func TestTryEnvs(t *testing.T) {
	os.Setenv("test_override_dir", "foo")

	result := tryEnvs([]string{"test_override_dir"}, "fallback_value")
	if result == nil {
		t.Errorf("Setting Environment failed")
	}
	if *result != "foo" {
		t.Errorf("Expected value to be 'foo', was '%s'", *result)
	}
}

func TestTryEnvsMultiValue(t *testing.T) {
	os.Setenv("test_override_dir2", "foo")

	result := tryEnvs([]string{"test_override_dir1", "test_override_dir2"}, "fallback_value")
	if result == nil {
		t.Errorf("Setting Environment failed")
	}
	if *result != "foo" {
		t.Errorf("Expected value to be 'foo', was '%s'", *result)
	}
}

func TestTryEnvsFallback(t *testing.T) {
	result := tryEnvs([]string{}, "fallback_value")
	if result == nil {
		t.Errorf("Setting Environment failed")
	}
	if *result != "fallback_value" {
		t.Errorf("Expected value to be 'fallback_value', was '%s'", *result)
	}
}

func TestGetApplicationDirectory(t *testing.T) {
	if runtime.GOOS == "linux" {
		// no implementation exists for linux
		return
	}

	result, err := GetApplicationDirectory()
	handleErr(err, t)
	handleResult(
		regexp.MustCompile("Program Files"),
		nil,
		regexp.MustCompile("Library/Application Support"),
		result,
		t,
	)
}

func TestGetUserConfigDirectory(t *testing.T) {
	result, err := GetUserConfigDirectory()
	handleErr(err, t)
	handleResult(
		regexp.MustCompile("AppData\\\\Roaming"),
		regexp.MustCompile("\\.config"),
		regexp.MustCompile("Library/Preferences"),
		result,
		t,
	)
}

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Got error %s", err)
	}
}

func handleResult(
	expectWindows *regexp.Regexp,
	expectLinux *regexp.Regexp,
	expectDarwin *regexp.Regexp,
	result string,
	t *testing.T,
) {
	switch runtime.GOOS {
	case "windows":
		if !expectWindows.MatchString(result) {
			t.Errorf("Result '%s' did not match expected expression '%s'", result, expectWindows.String())
		}
	case "linux":
		if !expectLinux.MatchString(result) {
			t.Errorf("Result '%s' did not match expected expression '%s'", result, expectLinux.String())
		}
	case "darwin":
		if !expectDarwin.MatchString(result) {
			t.Errorf("Result '%s' did not match expected expression '%s'", result, expectDarwin.String())
		}
	default:
		panic("Unsupported/unknown OS " + runtime.GOOS)
	}
}
