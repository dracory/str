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
