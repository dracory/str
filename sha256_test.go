package str_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/dracory/str"
)

func TestSHA256(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "EmptyString", input: ""},
		{name: "SimpleString", input: "hello"},
		{name: "SpecialChars", input: "hello!@#$%"},
		{name: "LongString", input: "The quick brown fox jumps over the lazy dog"},
		{name: "Unicode", input: "你好，世界！"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := sha256.Sum256([]byte(tt.input))
			result := str.SHA256(tt.input)
			expectedHex := fmt.Sprintf("%x", expected)

			if result != expectedHex {
				t.Errorf("SHA256(%q) = %q, want %q", tt.input, result, expectedHex)
			}
		})
	}
}
