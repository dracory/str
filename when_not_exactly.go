package str

// WhenNotExactly executes the truthy callback when the string does not exactly
// match the provided value; otherwise it executes the fallback if present.
func WhenNotExactly(str, value string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, !Exactly(str, value), truthy, fallback...)
}
