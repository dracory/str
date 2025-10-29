package str

// WhenContainsAll executes the truthy callback when the string contains all
// provided values. Otherwise it executes the fallback if present.
func WhenContainsAll(str string, values []string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, ContainsAll(str, values...), truthy, fallback...)
}
