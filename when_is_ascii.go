package str

// WhenIsAscii executes the truthy callback when the string is ASCII-only;
// otherwise it executes the fallback if supplied.
func WhenIsAscii(str string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, IsAscii(str), truthy, fallback...)
}
