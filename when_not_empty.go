package str

// WhenNotEmpty executes the truthy callback when the string is not empty;
// otherwise it executes the fallback if provided.
func WhenNotEmpty(str string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, IsNotEmpty(str), truthy, fallback...)
}
