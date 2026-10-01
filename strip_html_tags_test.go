package str_test

import (
	"testing"

	"github.com/dracory/str"
)

func TestStripHTMLTags(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic html tag removal",
			input:    "<p>Hello <b>World</b></p>",
			expected: "Hello World",
		},
		{
			name:     "script tags content removal",
			input:    "Hello <script>alert('xss')</script>World",
			expected: "Hello World",
		},
		{
			name:     "style tags content removal",
			input:    "Hello <style>body { color: red; }</style>World",
			expected: "Hello World",
		},
		{
			name:     "html entities decoding",
			input:    "Foo &amp; Bar &lt;Baz&gt; &quot;Qux&quot;",
			expected: "Foo & Bar <Baz> \"Qux\"",
		},
		{
			name:     "whitespace collapse and trim",
			input:    "  <p>  Multiple   spaces  </p>  ",
			expected: "Multiple spaces",
		},
		{
			name:     "plain text unmodified",
			input:    "Plain text without HTML",
			expected: "Plain text without HTML",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := str.StripHTMLTags(tt.input)
			if result != tt.expected {
				t.Errorf("StripHTMLTags(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
