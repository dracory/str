package str

// Unless executes the fallback when the predicate returns false, otherwise it
// returns the original string unchanged.
func Unless(str string, predicate func(string) bool, fallback func(string) string) string {
	if predicate == nil {
		return str
	}

	if predicate(str) {
		return str
	}

	if fallback != nil {
		return fallback(str)
	}

	return str
}
