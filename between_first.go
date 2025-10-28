package str

import "strings"

// BetweenFirst returns the substring between the first occurrence of the start string
// and the first occurrence of the end string after the start.
// If either start or end is empty or not found, returns empty string.
// Example: BetweenFirst("Hello [World] and [Universe]", "[", "]") returns "World"
func BetweenFirst(str, start, end string) string {
	if start == "" || end == "" {
		return ""
	}
	
	// Find start position
	startIdx := strings.Index(str, start)
	if startIdx == -1 {
		return ""
	}
	
	// Get substring after start
	afterStart := str[startIdx+len(start):]
	
	// Find end position in the substring after start
	endIdx := strings.Index(afterStart, end)
	if endIdx == -1 {
		return ""
	}
	
	return afterStart[:endIdx]
}
