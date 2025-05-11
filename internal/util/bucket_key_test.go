package util

import "testing"

func Test_ParseCompanyNameFromKey_OutputsCorrectNames(t *testing.T) {
	tests := []struct {
		key      string
		expected string
		hasError bool
	}{
		{"duckduckgo/rendered.html", "duckduckgo", false},
		{"posthog/snapshot.html", "posthog", false},
		{"", "", true},
		{"/weird.html", "", true},
	}

	for _, tt := range tests {
		result, err := ParseCompanyNameFromKey(tt.key)
		if tt.hasError && err == nil {
			t.Errorf("expected error for key %q, got none", tt.key)
		}
		if !tt.hasError && result != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, result)
		}
	}
}
