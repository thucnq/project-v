//go:build freebsd
// +build freebsd

package idgen

import (
	"bytes"
	"io/ioutil"
	"os"
	"syscall"
)

const hostidPath = "/etc/hostid"

// machineID returns the uuid specified at `/etc/hostid`.
// If the returned value is empty, the uuid from a call to `kenv -q smbios.system.uuid` is returned.
// If there is an error an empty string is returned.
func readPlatformMachineID() (string, error) {
	id, err := readHostid()
	if err != nil {
		// try fallback
		id, err = readKenv()
	}
	if err != nil {
		// try fallback
		id, err = readHostuuid()
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func readHostid() (string, error) {
	buf, err := ioutil.ReadFile(hostidPath)
	if err != nil {
		return "", err
	}
	return trim(string(buf)), nil
}

func readKenv() (string, error) {
	buf := &bytes.Buffer{}
	err := run(buf, os.Stderr, "kenv", "-q", "smbios.system.uuid")
	if err != nil {
		return "", err
	}
	return trim(buf.String()), nil
}

func readHostuuid() (string, error) {
	return syscall.Sysctl("kern.hostuuid")
}
