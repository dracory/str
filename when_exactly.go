package str

// WhenExactly executes the truthy callback when the string exactly matches the
// provided value; otherwise it executes the fallback if present.
func WhenExactly(str, value string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, Exactly(str, value), truthy, fallback...)
}
