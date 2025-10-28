package str

// import (
// 	"bytes"
// 	"crypto/rand"
// 	"encoding/json"
// 	"path/filepath"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"unicode"
// 	"unicode/utf8"

// 	"golang.org/x/exp/constraints"
// 	"golang.org/x/text/cases"
// 	"golang.org/x/text/language"
// )

// type String struct {
// 	value string
// }

// // ExcerptOption is the option for Excerpt method
// type ExcerptOption struct {
// 	Radius   int
// 	Omission string
// }
// // Of creates a new String instance with the given value.
// func Of(value string) *String {
// 	return &String{value: value}
// }

// // Finish returns the String instance with the given value appended.
// // If the given value already ends with the suffix, it will not be added twice.
// func (s *String) Finish(value string) *String {
// 	quoted := regexp.QuoteMeta(value)
// 	reg := regexp.MustCompile("(?:" + quoted + ")+$")
// 	s.value = reg.ReplaceAllString(s.value, "") + value
// 	return s
// }

// // Kebab returns the String instance in kebab case.
// func (s *String) Kebab() *String {
// 	return s.Snake("-")
// }

// // Limit returns the String instance truncated to the given length.
// func (s *String) Limit(limit int, end ...string) *String {
// 	defaultEnd := "..."
// 	if len(end) > 0 {
// 		defaultEnd = end[0]
// 	}

// 	if s.Length() <= limit {
// 		return s
// 	}
// 	s.value = Substr(s.value, 0, limit) + defaultEnd
// 	return s
// }

// // Lower returns the String instance in lower case.
// func (s *String) Lower() *String {
// 	s.value = strings.ToLower(s.value)
// 	return s
// }

// // Ltrim returns the String instance with the leftmost occurrence of the given value removed.
// func (s *String) LTrim(characters ...string) *String {
// 	if len(characters) == 0 {
// 		s.value = strings.TrimLeft(s.value, " ")
// 		return s
// 	}

// 	s.value = strings.TrimLeft(s.value, characters[0])
// 	return s
// }
// // Test returns true if the string matches the given pattern.
// func (s *String) Test(pattern string) bool {
// 	return s.IsMatch(pattern)
// }

// // Title returns the String instance in title case.
// func (s *String) Title() *String {
// 	casesTitle := cases.Title(language.Und)
// 	s.value = casesTitle.String(s.value)
// 	return s
// }

// // Trim returns the String instance with trimmed characters from the left and right sides.
// func (s *String) Trim(characters ...string) *String {
// 	if len(characters) == 0 {
// 		s.value = strings.TrimSpace(s.value)
// 		return s
// 	}

// 	s.value = strings.Trim(s.value, characters[0])
// 	return s
// }

// // Unless returns the String instance with the given fallback applied if the given condition is false.
// func (s *String) Unless(callback func(*String) bool, fallback func(*String) *String) *String {
// 	if !callback(s) {
// 		return fallback(s)
// 	}

// 	return s
// }

// // When returns the String instance with the given callback applied if the given condition is true.
// // If the condition is false, the fallback callback is applied (if provided).
// func (s *String) When(condition bool, callback ...func(*String) *String) *String {
// 	if condition {
// 		if len(callback) > 0 && callback[0] != nil {
// 			return callback[0](s)
// 		}
// 	} else {
// 		if len(callback) > 1 && callback[1] != nil {
// 			return callback[1](s)
// 		}
// 	}

// 	return s
// }

// // WhenContains returns the String instance with the given callback applied if the string contains the given value.
// func (s *String) WhenContains(value string, callback ...func(*String) *String) *String {
// 	return s.When(s.Contains(value), callback...)
// }

// // WhenContainsAll returns the String instance with the given callback applied if the string contains all the given values.
// func (s *String) WhenContainsAll(values []string, callback ...func(*String) *String) *String {
// 	return s.When(s.ContainsAll(values...), callback...)
// }

// // WhenEmpty returns the String instance with the given callback applied if the string is empty.
// func (s *String) WhenEmpty(callback ...func(*String) *String) *String {
// 	return s.When(s.IsEmpty(), callback...)
// }

// // WhenIsAscii returns the String instance with the given callback applied if the string contains only ASCII characters.
// func (s *String) WhenIsAscii(callback ...func(*String) *String) *String {
// 	return s.When(s.IsAscii(), callback...)
// }

// // WhenNotEmpty returns the String instance with the given callback applied if the string is not empty.
// func (s *String) WhenNotEmpty(callback ...func(*String) *String) *String {
// 	return s.When(s.IsNotEmpty(), callback...)
// }

// // WhenStartsWith returns the String instance with the given callback applied if the string starts with the given value.
// func (s *String) WhenStartsWith(value []string, callback ...func(*String) *String) *String {
// 	return s.When(s.StartsWith(value...), callback...)
// }

// // WhenEndsWith returns the String instance with the given callback applied if the string ends with the given value.
// func (s *String) WhenEndsWith(value []string, callback ...func(*String) *String) *String {
// 	return s.When(s.EndsWith(value...), callback...)
// }

// // WhenExactly returns the String instance with the given callback applied if the string is exactly the given value.
// func (s *String) WhenExactly(value string, callback ...func(*String) *String) *String {
// 	return s.When(s.Exactly(value), callback...)
// }

// // WhenNotExactly returns the String instance with the given callback applied if the string is not exactly the given value.
// func (s *String) WhenNotExactly(value string, callback ...func(*String) *String) *String {
// 	return s.When(!s.Exactly(value), callback...)
// }

// // WhenIs returns the String instance with the given callback applied if the string matches any of the given patterns.
// func (s *String) WhenIs(value string, callback ...func(*String) *String) *String {
// 	return s.When(s.Is(value), callback...)
// }

// // WhenIsUlid returns the String instance with the given callback applied if the string is a valid ULID.
// func (s *String) WhenIsUlid(callback ...func(*String) *String) *String {
// 	return s.When(s.IsUlid(), callback...)
// }

// // WhenIsUuid returns the String instance with the given callback applied if the string is a valid UUID.
// func (s *String) WhenIsUuid(callback ...func(*String) *String) *String {
// 	return s.When(s.IsUuid(), callback...)
// }

// // WhenTest returns the String instance with the given callback applied if the string matches the given pattern.
// func (s *String) WhenTest(pattern string, callback ...func(*String) *String) *String {
// 	return s.When(s.Test(pattern), callback...)
// }

// type Buffer struct {
// 	*bytes.Buffer
// }

// func NewBuffer() *Buffer {
// 	return &Buffer{Buffer: new(bytes.Buffer)}
// }

// func (b *Buffer) Append(i any) *Buffer {
// 	switch val := i.(type) {
// 	case int:
// 		b.append(strconv.Itoa(val))
// 	case int64:
// 		b.append(strconv.FormatInt(val, 10))
// 	case uint:
// 		b.append(strconv.FormatUint(uint64(val), 10))
// 	case uint64:
// 		b.append(strconv.FormatUint(val, 10))
// 	case string:
// 		b.append(val)
// 	case []byte:
// 		b.Write(val)
// 	case rune:
// 		b.WriteRune(val)
// 	}
// 	return b
// }

// func (b *Buffer) append(s string) *Buffer {
// 	b.WriteString(s)

// 	return b
// }

// // maximum returns the largest of x or y.
// func maximum[T constraints.Ordered](x T, y T) T {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
