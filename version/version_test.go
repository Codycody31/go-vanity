package version

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "Version is set",
			expected: "1.0.0",
		},
		{
			name:     "Version is not set",
			expected: "dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.expected
			actual := String()
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
