package utils

import (
	"regexp"
	"strings"
)

// Function to determine if the delta between two versions is either a major, minor or patch change
//
// Parameters:
//   - remote_version: The remote version to compare
//   - local_version: The local version to compare
//
// Returns:
//   - delta: The delta between the two versions or "unknown" if the delta cannot be determined, none if version are the same
func ComputeDelta(remote_version string, local_version string) string {

	// Check if version are well formed
	// Follow regex: ^v?[0-9]+\.[0-9]+\.[0-9]+$
	var re = regexp.MustCompile(`(?mi)^v?[0-9]+\.[0-9]+\.[0-9]+$`)
	if !re.MatchString(remote_version) || !re.MatchString(local_version) {
		return "unknown"
	}

	// Remove the v prefix if present
	remote_version = strings.TrimPrefix(remote_version, "v")
	local_version = strings.TrimPrefix(local_version, "v")

	if remote_version == local_version {
		return "none"
	}

	// Split remote version with dots
	remote_version_split := strings.Split(remote_version, ".")
	// Split local version with dots
	local_version_split := strings.Split(local_version, ".")

	// if remote version is greater than local version
	if remote_version_split[0] > local_version_split[0] {
		return "major"
	} else if remote_version_split[0] == local_version_split[0] {
		if remote_version_split[1] > local_version_split[1] {
			return "minor"
		} else if remote_version_split[1] == local_version_split[1] {
			if remote_version_split[2] > local_version_split[2] {
				return "patch"
			}
		}
	}

	return "unknown"
}
