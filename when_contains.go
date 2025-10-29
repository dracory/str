package str

// WhenContains executes the truthy callback when the string contains value.
// Otherwise it executes the fallback if provided.
func WhenContains(str, value string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, Contains(str, value), truthy, fallback...)
}
