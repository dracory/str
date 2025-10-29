package str

// WhenIsUuid executes the truthy callback when the string is a valid UUID;
// otherwise it executes the fallback if present.
func WhenIsUuid(str string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, IsUuid(str), truthy, fallback...)
}
