package pkg

import (
	"fmt"
	"strings"
)

const (
	MinimumSDKVersion = "1.0.0"
	MaximumSDKVersion = ""
)

// Version represents a semantic version
type Version struct {
	Major, Minor, Patch int
}

// ParseVersion parses a version string into a Version struct
func ParseVersion(version string) (*Version, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid version format: %s", version)
	}

	var v Version
	_, err := fmt.Sscanf(parts[0], "%d", &v.Major)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Sscanf(parts[1], "%d", &v.Minor)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Sscanf(parts[2], "%d", &v.Patch)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// Compare compares two versions
func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		return v.Major - other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor - other.Minor
	}
	return v.Patch - other.Patch
}

// EnsureCompatibility checks if a version meets minimum and maximum requirements
func EnsureCompatibility(version, minVersion, maxVersion string) error {
	v, err := ParseVersion(version)
	if err != nil {
		return fmt.Errorf("invalid version format: %v", err)
	}

	if minVersion != "" {
		min, err := ParseVersion(minVersion)
		if err != nil {
			return fmt.Errorf("invalid minimum version format: %v", err)
		}
		if v.Compare(min) < 0 {
			return fmt.Errorf("version %s is smaller than minimum version %s", version, minVersion)
		}
	}

	if maxVersion != "" {
		max, err := ParseVersion(maxVersion)
		if err != nil {
			return fmt.Errorf("invalid maximum version format: %v", err)
		}
		if v.Compare(max) > 0 {
			return fmt.Errorf("version %s is larger than maximum version %s", version, maxVersion)
		}
	}

	return nil
} 