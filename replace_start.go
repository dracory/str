package str

// ReplaceStart replaces the beginning of the string if it starts with search.
// If search is empty or not present at the start, the original string is returned.
func ReplaceStart(str, search, replace string) string {
	if search == "" {
		return str
	}

	if len(str) < len(search) {
		return str
	}

	if len(str) >= len(search) && str[:len(search)] == search {
		return replace + str[len(search):]
	}

	return str
}
