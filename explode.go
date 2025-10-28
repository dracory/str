package str

import "strings"

// Explode splits the input string using the provided delimiter.
// An optional limit mirrors PHP's explode semantics:
//   >0 keeps at most limit elements, combining the remainder into the last element.
//   <0 omits the last |limit| elements.
//   0 is treated as if no limit was provided.
func Explode(str, delimiter string, limit ...int) []string {
	defaultLimit := 1
	isLimitSet := false
	if len(limit) > 0 && limit[0] != 0 {
		defaultLimit = limit[0]
		isLimitSet = true
	}

	parts := strings.Split(str, delimiter)

	if !isLimitSet || len(parts) <= defaultLimit {
		return parts
	}

	if defaultLimit > 0 {
		head := append([]string{}, parts[:defaultLimit-1]...)
		head = append(head, strings.Join(parts[defaultLimit-1:], delimiter))
		return head
	}

	if defaultLimit < 0 && len(parts) <= -defaultLimit {
		return []string{}
	}

	return parts[:len(parts)+defaultLimit]
}
