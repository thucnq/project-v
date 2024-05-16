//go:build darwin
// +build darwin

package idgen

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// run wraps `exec.Command` with easy access to stdout and stderr.
func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = stdout
	c.Stderr = stderr
	return c.Run()
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}

func readPlatformMachineID() (string, error) {
	id, err := syscall.Sysctl("kern.uuid")
	if err != nil {
		buf := &bytes.Buffer{}
		err = run(buf, os.Stderr, "ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
		if err != nil {
			return "", err
		}
		id, err = extractID(buf.String())
		if err != nil {
			return "", err
		}
		return trim(id), nil
	}
	return id, nil
}
func extractID(lines string) (string, error) {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `"`), nil
			}
		}
	}
	return "", fmt.Errorf("Failed to extract 'IOPlatformUUID' value from `ioreg` output.\n%s", lines)
}
