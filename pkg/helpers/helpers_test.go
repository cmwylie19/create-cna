package helpers

import (
	"testing"
)

func TestByteReplace(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    []byte
		old      string
		new      string
		expected []byte
	}{
		{
			name:     "replace single occurrence",
			input:    []byte("hello, world!"),
			old:      "world",
			new:      "go",
			expected: []byte("hello, go!"),
		},
		{
			name:     "replace multiple occurrences",
			input:    []byte("the quick brown fox jumped over the lazy dog"),
			old:      "the",
			new:      "a",
			expected: []byte("a quick brown fox jumped over a lazy dog"),
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call ByteReplace with test case input
			output := ByteReplace(tc.input, tc.old, tc.new)

			// Compare actual output with expected output
			if string(output) != string(tc.expected) {
				t.Errorf("expected '%s', but got '%s'", tc.expected, output)
			}
		})
	}
}
