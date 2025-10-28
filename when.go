package str

// When executes the truthy callback when condition is true, otherwise executes
// the fallback if provided. It returns the resulting string.
func When(str string, condition bool, truthy func(string) string, fallback ...func(string) string) string {
	if condition {
		if truthy != nil {
			return truthy(str)
		}
		return str
	}

	if len(fallback) > 0 && fallback[0] != nil {
		return fallback[0](str)
	}

	return str
}
