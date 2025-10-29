package str

// WhenIsUlid executes the truthy callback when the string is a valid ULID;
// otherwise it executes the fallback if present.
func WhenIsUlid(str string, truthy func(string) string, fallback ...func(string) string) string {
	return When(str, IsUlid(str), truthy, fallback...)
}
