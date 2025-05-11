package util

import (
	"fmt"
	"strings"
)

// ParseCompanyNameFromKey extracts the company name from an S3 key of the form "company/somefile.html"
func ParseCompanyNameFromKey(key string) (string, error) {
	segments := strings.SplitN(key, "/", 2)
	if len(segments) < 1 || segments[0] == "" {
		return "", fmt.Errorf("invalid key format: %s", key)
	}
	return segments[0], nil
}
