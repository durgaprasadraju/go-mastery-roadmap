package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// SemVer holds parsed version.
type SemVer struct{ Major, Minor, Patch int }

// ParseSemVer parses major.minor.patch.
func ParseSemVer(s string) (SemVer, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return SemVer{}, fmt.Errorf("invalid semver")
	}
	maj, err := strconv.Atoi(parts[0])
	if err != nil {
		return SemVer{}, err
	}
	min, err := strconv.Atoi(parts[1])
	if err != nil {
		return SemVer{}, err
	}
	pat, err := strconv.Atoi(parts[2])
	if err != nil {
		return SemVer{}, err
	}
	return SemVer{maj, min, pat}, nil
}