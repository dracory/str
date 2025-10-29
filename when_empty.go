package str

// WhenEmpty executes the truthy callback when the string is empty; otherwise it
// executes the fallback if supplied.
func WhenEmpty(str string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, IsEmpty(str), truthy, fallback...)
}
