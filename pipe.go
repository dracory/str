package str

// Pipe passes the string through the provided callback and returns the result.
// If callback is nil, the original string is returned unchanged.
func Pipe(in string, callback func(string) string) string {
	if callback == nil {
		return in
	}

	return callback(in)
}
