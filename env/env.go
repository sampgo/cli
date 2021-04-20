package env

import (
	"os"
	"runtime"
)

// Set configures our cgo related environment variables.
func Set() {
	os.Setenv("CGO_CFLAGS_ALLOW", ".*")
	os.Setenv("CGO_LDFLAGS_ALLOW", ".*")

	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("GOARCH", "386")

	if runtime.GOOS == "windows" {
		// set GOOS to windows on windows.
		os.Setenv("GOOS", "windows")
	} else {
		// any other OS? set it to linux.
		os.Setenv("GOOS", "linux")
	}
}

// Unset clears the previously set cgo related environment variables.
func Unset() {
	os.Unsetenv("CGO_CFLAGS_ALLOW")
	os.Unsetenv("CGO_LDFLAGS_ALLOW")

	os.Unsetenv("CGO_ENABLED")
	os.Unsetenv("GOARCH")

	os.Unsetenv("GOOS")
}
