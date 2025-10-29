package str

// Test returns true if the string matches the given regular expression pattern.
// Empty patterns always return false.
func Test(str, pattern string) bool {
	if pattern == "" {
		return false
	}

	return IsMatch(str, pattern)
}
