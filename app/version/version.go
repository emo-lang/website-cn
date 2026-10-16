// Package version carries the site version string. main reads it from the
// VERSION file at boot and records it here, so page templates render the same
// value that `--version` prints.
package version

var value string

// Set records the version read from the VERSION file at boot.
func Set(v string) {
	value = v
}

// String returns the site version, e.g. "v0.26.11".
func String() string {
	return value
}
