//go:build linux
// +build linux

package idgen

import (
	"io/ioutil"
	"strings"
)

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"
)

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}
func readPlatformMachineID() (string, error) {
	id, err := ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	if err != nil {
		// try fallback path
		id, err = ioutil.ReadFile(dbusPath)
	}
	if err != nil {
		// try fallback path
		id, err = ioutil.ReadFile(dbusPathEtc)
	}
	if err != nil {
		return "", err
	}
	return trim(string(id)), nil
}
