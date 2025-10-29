package str

// WhenIs executes the truthy callback when the string matches any of the
// provided glob patterns; otherwise it executes the fallback if present.
func WhenIs(str string, patterns []string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, Is(str, patterns...), truthy, fallback...)
}
