// +build !linux

package userns

import (
	"os"

	"github.com/opencontainers/runc/libcontainer/user"
)

// runningInUserNS is a stub for non-Linux systems
// Always returns false
func runningInUserNS() bool {
	return false
}

// uidMapInUserNS is a stub for non-Linux systems
// Always returns false
func uidMapInUserNS(uidmap []user.IDMap) bool {
	return false
}

// getParentNSeuid returns the euid within the parent user namespace
// Always returns os.Geteuid on non-linux
func getParentNSeuid() int {
	return os.Geteuid()
}
