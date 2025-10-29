package str

// WhenStartsWith executes the truthy callback when the string starts with any
// of the provided values; otherwise it executes the fallback if present.
func WhenStartsWith(str string, values []string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, StartsWith(str, values...), truthy, fallback...)
}
