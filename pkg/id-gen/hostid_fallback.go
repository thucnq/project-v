//go:build !darwin && !linux && !freebsd && !windows
// +build !darwin,!linux,!freebsd,!windows

package idgen

import "errors"

func readPlatformMachineID() (string, error) {
	return "", errors.New("not implemented")
}
