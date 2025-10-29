package str

import (
	"bytes"
	"strconv"
)

// buffer wraps bytes.Buffer providing fluent append helpers for primitive types.
type buffer struct {
	*bytes.Buffer
}

// NewBuffer creates a new buffer instance.
func NewBuffer() *buffer {
	return &buffer{Buffer: new(bytes.Buffer)}
}

// Append writes the value to the buffer, converting primitive types to strings
// where appropriate. It returns the buffer to support chaining.
func (b *buffer) Append(v any) *buffer {
	switch val := v.(type) {
	case int:
		return b.appendString(strconv.Itoa(val))
	case int64:
		return b.appendString(strconv.FormatInt(val, 10))
	case uint:
		return b.appendString(strconv.FormatUint(uint64(val), 10))
	case uint64:
		return b.appendString(strconv.FormatUint(val, 10))
	case string:
		return b.appendString(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	case bool:
		return b.appendString(strconv.FormatBool(val))
	}

	return b
}

func (b *buffer) appendString(s string) *buffer {
	b.WriteString(s)
	return b
}
