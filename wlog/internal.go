package wlog

import (
	"runtime"
	"strings"
)

// donot ask for other package

// IsWindows judge host os system.
func IsWindows() bool {
	t := strings.ToUpper(runtime.GOOS)
	return strings.Contains(t, "WINDOW")
}

func expandPath(path string) string {
	return path
}
