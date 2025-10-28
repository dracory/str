package str

import (
	"strings"
	"unicode"
)

type StringBuilder struct {
	value string
}

func Of(value string) *StringBuilder {
	return &StringBuilder{value: value}
}

func (sb *StringBuilder) LTrim(chars ...string) *StringBuilder {
	if len(chars) == 0 {
		sb.value = strings.TrimLeftFunc(sb.value, unicode.IsSpace)
		return sb
	}

	if chars[0] == "" {
		sb.value = strings.TrimLeftFunc(sb.value, unicode.IsSpace)
		return sb
	}

	sb.value = strings.TrimLeft(sb.value, chars[0])
	return sb
}

func (sb *StringBuilder) RTrim(chars ...string) *StringBuilder {
	if len(chars) == 0 {
		sb.value = strings.TrimRightFunc(sb.value, unicode.IsSpace)
		return sb
	}

	if chars[0] == "" {
		sb.value = strings.TrimRightFunc(sb.value, unicode.IsSpace)
		return sb
	}

	sb.value = strings.TrimRight(sb.value, chars[0])
	return sb
}

func (sb *StringBuilder) Append(parts ...string) *StringBuilder {
	sb.value += strings.Join(parts, "")
	return sb
}

func (sb *StringBuilder) Prepend(prefix string) *StringBuilder {
	sb.value = prefix + sb.value
	return sb
}

func (sb *StringBuilder) Exactly(str string) bool {
	return sb.value == str
}

func (sb *StringBuilder) Unless(predicate func(*StringBuilder) bool, fallback func(*StringBuilder) *StringBuilder) *StringBuilder {
	if predicate(sb) {
		return sb
	}

	return fallback(sb)
}

func (sb *StringBuilder) String() string {
	return sb.value
}
