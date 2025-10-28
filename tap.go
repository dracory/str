package str

// Tap passes the string to the provided callback and returns the original
// string unchanged. A nil callback is safely ignored.
func Tap(str string, callback func(string)) string {
	if callback != nil {
		callback(str)
	}

	return str
}
