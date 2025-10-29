package str

// Kebab converts the input string to kebab-case.
func Kebab(str string) string {
	return Snake(str, "-")
}
