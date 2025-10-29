package str

// WhenEndsWith executes the truthy callback when the string ends with any of
// the provided values; otherwise it executes the fallback if present.
func WhenEndsWith(str string, values []string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, EndsWith(str, values...), truthy, fallback...)
}
