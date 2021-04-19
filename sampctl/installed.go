package sampctl

import "os/exec"

// Installed returns sampctl's installation status on this system.
func Installed() bool {
	_, err := exec.LookPath("sampctl")

	return err == nil
}
