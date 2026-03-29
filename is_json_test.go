package str_test

import (
	"testing"

	"github.com/dracory/str"
)

func TestIsJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Valid JSON Object",
			input:    `{"key": "value"}`,
			expected: true,
		},
		{
			name:     "Valid JSON Array",
			input:    `[1, 2, 3]`,
			expected: true,
		},
		{
			name:     "Empty JSON Object",
			input:    `{}`,
			expected: true,
		},
		{
			name:     "Empty JSON Array",
			input:    `[]`,
			expected: true,
		},
		{
			name:     "Invalid - Only Opening Brace",
			input:    `{"key": "value"`,
			expected: false,
		},
		{
			name:     "Invalid - Only Closing Brace",
			input:    `"key": "value"}`,
			expected: false,
		},
		{
			name:     "Invalid - Only Opening Bracket",
			input:    `[1, 2, 3`,
			expected: false,
		},
		{
			name:     "Invalid - Only Closing Bracket",
			input:    `1, 2, 3]`,
			expected: false,
		},
		{
			name:     "Invalid - Reversed Braces",
			input:    `}{`,
			expected: false,
		},
		{
			name:     "Invalid - Reversed Brackets",
			input:    `][`,
			expected: false,
		},
		{
			name:     "Plain String",
			input:    "hello",
			expected: false,
		},
		{
			name:     "JSON String Value",
			input:    `"just a string"`,
			expected: false,
		},
		{
			name:     "Whitespace Only",
			input:    "   ",
			expected: false,
		},
		{
			name:     "Nested JSON Object",
			input:    `{"outer": {"inner": "value"}}`,
			expected: true,
		},
		{
			name:     "Nested JSON Array",
			input:    `[[1, 2], [3, 4]]`,
			expected: true,
		},
		{
			name:     "JSON Object with Whitespace",
			input:    `  {"key": "value"}  `,
			expected: false,
		},
		{
			name:     "JSON Array with Whitespace",
			input:    `  [1, 2, 3]  `,
			expected: false,
		},
		{
			name:     "Invalid - XML-like",
			input:    `<tag>value</tag>`,
			expected: false,
		},
		{
			name:     "Invalid - Number",
			input:    "12345",
			expected: false,
		},
		{
			name:     "Invalid - Boolean",
			input:    "true",
			expected: false,
		},
		{
			name:     "Invalid - Null",
			input:    "null",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsJSON(tc.input)
			if result != tc.expected {
				t.Errorf("IsJSON(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
