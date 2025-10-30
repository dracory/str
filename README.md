# Str Package

[![Tests Status](https://github.com/dracory/str/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/str/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/str)](https://goreportcard.com/report/github.com/dracory/str)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/str)](https://pkg.go.dev/github.com/dracory/str)

The `str` package provides a comprehensive set of string manipulation utilities for the Dracory framework.
It offers a wide range of functions for string operations, validation, transformation, and formatting.

## Overview

This package includes utilities for:

1. **String Manipulation**: Functions for modifying, extracting, and transforming strings
2. **String Validation**: Functions for checking string properties and patterns
3. **String Formatting**: Functions for formatting strings in various ways
4. **String Conversion**: Functions for converting between different string formats and encodings
5. **String Generation**: Functions for generating random strings and hashes
6. **Price Formatting**: Functions for formatting prices with currency symbols

## Key Features

- **Pattern Matching**: Check if strings match patterns using glob syntax
- **String Extraction**: Extract substrings based on various criteria
- **String Transformation**: Convert strings to different formats (camelCase, snake_case, etc.)
- **String Validation**: Check if strings are empty, UUIDs, ULIDs, etc.
- **String Encoding**: Encode and decode strings using various encodings (Base32, Base64, etc.)
- **String Hashing**: Generate hashes from strings (MD5, BCrypt, etc.)
- **String Formatting**: Format strings with padding, truncation, etc.
- **Conditional Helpers**: Run callbacks when predicates (empty, pattern match, UUID/ULID, etc.) succeed
- **String Generation**: Generate random strings and identifiers
- **Price Formatting**: Format prices with currency symbols (USD, GBP, EUR)

## Usage Examples

### String Validation

```go
import "github.com/dracory/str"

// Check if a string is empty
isEmpty := str.IsEmpty("")  // true

// Check if a string is not empty
isNotEmpty := str.IsNotEmpty("Hello")  // true

// Check if a string is a UUID
isUUID := str.IsUUID("123e4567-e89b-12d3-a456-426614174000")  // true

// Check if a string is a ULID
isULID := str.IsULID("01H9Z8K2P3M4N5Q6R7S8T9U0V")  // true

// Check if a string matches a pattern
matches := str.Is("hello.txt", "*.txt")  // true

// Check if a string matches a regular expression pattern
matchesRegex := str.Test("hello123", `^hello\d+$`)  // true
```

### String Manipulation

```go
import "github.com/dracory/str"

// Extract substring between two strings
between := str.Between("Hello [World] Test", "[", "]")  // "World"

// Extract substring before a string
before := str.Before("Hello World", "World")  // "Hello "

// Extract substring after a string
after := str.After("Hello World", "Hello ")  // "World"

// Extract substring before the last occurrence of a string
beforeLast := str.BeforeLast("Hello World World", "World")  // "Hello "

// Extract substring after the last occurrence of a string
afterLast := str.AfterLast("Hello World World", "World")  // ""

// Extract substring from the left
leftFrom := str.LeftFrom("Hello World", 5)  // "Hello"

// Extract substring from the right
rightFrom := str.RightFrom("Hello World", 5)  // "World"

// Truncate a string
truncated := str.Truncate("Hello World", 8, "...")  // "Hello..."

// Limit a string by rune length with optional suffix
limited := str.Limit("Hello World", 5)  // "Hello..."

// Convert to snake_case
snake := str.ToSnake("HelloWorld")  // "hello_world"

// Convert to kebab-case
kebab := str.Kebab("HelloWorld")  // "hello-world"

// Convert to camelCase
camel := str.ToCamel("hello_world")  // "helloWorld"

// Convert first character to uppercase
ucFirst := str.UcFirst("hello")  // "Hello"

// Convert string to uppercase
upper := str.Upper("hello")  // "HELLO"

// Split string into words
words := str.Words("Hello World")  // ["Hello", "World"]

// Count words in a string
wordCount := str.WordCount("Hello World")  // 2

// Create a URL-friendly slug
slug := str.Slugify("Hello World!", '-')  // "hello-world"
```

### String Encoding and Hashing

```go
import "github.com/dracory/str"

// Encode string to Base64
base64 := str.Base64Encode("Hello World")  // "SGVsbG8gV29ybGQ="

// Decode Base64 string
decoded := str.Base64Decode("SGVsbG8gV29ybGQ=")  // "Hello World"

// Encode string to Base32 Extended
base32 := str.Base32ExtendedEncode("Hello World")  // "91IMOR3FCPBI41"

// Decode Base32 Extended string
decoded := str.Base32ExtendedDecode("91IMOR3FCPBI41")  // "Hello World"

// Generate MD5 hash (legacy helper — DO NOT use for security-sensitive hashing)
md5 := str.MD5("Hello World")  // "b10a8db164e0754105b7a99be72e3fe5"

// Generate SHA-256 hash (preferred general-purpose hashing helper)
sha256 := str.SHA256("Hello World")  // "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"

// Generate BCrypt hash
bcrypt := str.ToBcryptHash("password")  // "$2a$10$..."

// Compare password with BCrypt hash
matches := str.BcryptHashCompare("password", bcrypt)  // true
```

### String Generation

```go
import "github.com/dracory/str"

// Generate a random string
random := str.Random(10)  // "a1b2c3d4e5"

// Generate a random string from a gamma distribution
randomGamma := str.RandomFromGamma(10, 2.0, 1.0)  // "a1b2c3d4e5"

// Convert integer to Base32
base32 := str.IntToBase32(12345)  // "3RP"

// Convert integer to Base36
base36 := str.IntToBase36(12345)  // "9IX"
```

### Conditional Helpers

```go
import "github.com/dracory/str"

result := str.When("admin", str.Is("admin", "admin"), func(s string) string {
    return s + "@example.com"
})
// result == "admin@example.com"

result = str.WhenEmpty("", func(_ string) string { return "fallback" })
// result == "fallback"

result = str.WhenIsUUID("123e4567-e89b-12d3-a456-426614174000", func(s string) string {
    return "valid:" + s
})
// result == "valid:123e4567-e89b-12d3-a456-426614174000"
```

### Buffer Utility

```go
import "github.com/dracory/str"

buf := str.NewBuffer()
buf.Append("user-")
buf.Append(42)
buf.Append(true)

result := buf.String()  // "user-42true"
```

### Price Formatting

```go
import "github.com/dracory/str"

// Get currency symbol (Unicode by default)
symbol := str.CurrencySymbol("USD")  // "$"
symbol := str.CurrencySymbol("GBP")  // "£"
symbol := str.CurrencySymbol("EUR")  // "€"
symbol := str.CurrencySymbol("JPY")  // "¥"
symbol := str.CurrencySymbol("INR")  // "₹"

// Get currency symbol as HTML entity
symbol := str.CurrencySymbol("GBP", true)  // "&pound;"
symbol := str.CurrencySymbol("EUR", true)  // "&euro;"

// Supports 40+ currencies including:
// Major: USD, EUR, GBP, JPY, CNY, INR, KRW, RUB, TRY
// Dollar variants: AUD, CAD, HKD, SGD, NZD, MXN, BRL, ARS
// European: CHF, SEK, NOK, DKK, PLN, CZK, HUF, RON, BGN
// Middle East & Africa: SAR, AED, ILS, ZAR, EGP, NGN, KES
// Asian: THB, IDR, MYR, PHP, VND, PKR, BDT
// Crypto: BTC, ETH

// Convert float to formatted price string (Unicode by default)
price := str.ToPrice(19.99, "USD")  // "$19.99"
price := str.ToPrice(100.00, "GBP")  // "£100.00"
price := str.ToPrice(5000.50, "JPY")  // "¥5000.50"

// Convert float to formatted price string with HTML entity
price := str.ToPrice(19.99, "GBP", true)  // "&pound;19.99"
price := str.ToPrice(45.50, "EUR", true)  // "&euro;45.50"

// Convert string to formatted price string (with error handling)
price, err := str.ToPriceFromString("19.99", "USD")  // "$19.99", nil
price, err := str.ToPriceFromString("invalid", "USD")  // "", error

// Convert string to formatted price string with HTML entity
price, err := str.ToPriceFromString("19.99", "GBP", true)  // "&pound;19.99", nil

// Convert string to formatted price string (with default fallback)
price := str.ToPriceFromStringOrDefault("19.99", "USD", "n/a")  // "$19.99"
price := str.ToPriceFromStringOrDefault("invalid", "USD", "n/a")  // "n/a"

// Convert string to formatted price string with HTML entity and default fallback
price := str.ToPriceFromStringOrDefault("19.99", "EUR", "n/a", true)  // "&euro;19.99"

## Function Reference

The following lists every exported helper, grouped by their primary area. Functions that accept variadic arguments support optional parameters for fine-grained control.

### Validation & Matching

- `Is(str, patterns ...string)` — checks glob patterns (e.g., `*.txt`).
- `IsMatch(str, pattern string)` — tests strings against regular expressions.
- `Test(str, pattern string)` — convenience regex check that ignores empty patterns.
- `Contains(str, value string)` — verifies a substring exists.
- `ContainsAll(str string, values []string)` — ensures all values occur in the string.
- `ContainsAnyChar(str, chars string)` — reports whether any rune from `chars` is present.
- `ContainsOnly(str, chars string)` — verifies every rune is within `chars`.
- `StartsWith(str string, prefixes ...string)` — checks prefix matches.
- `EndsWith(str string, suffixes ...string)` — checks suffix matches.
- `Exactly(str, value string)` — requires an exact match (string function); `(*StringBuilder).Exactly` compares the builder value.
- `Match(str, pattern string)` — finds the first regex match and returns it.
- `MatchAll(str, pattern string)` — returns all regex matches.
- `IsEmpty(str string)` / `IsNotEmpty(str string)` — emptiness predicates.
- `IsAscii(str string)` — ensures the string is ASCII-only.
- `IsUuid(str string)` / `IsUlid(str string)` — identifier validation helpers.
- `IsMap(v any)` / `IsSlice(v any)` — type assertions for dynamic values.

### Extraction & Structure

- `After(str, search string)` / `AfterLast(str, search string)` — substrings after a delimiter.
- `Before(str, search string)` / `BeforeLast(str, search string)` — substrings before a delimiter.
- `Between(str, start, end string)` / `BetweenFirst(str, start, end string)` — segments between markers.
- `LeftFrom(str string, index int)` / `RightFrom(str string, index int)` — slice from rune offsets.
- `Substr(str string, offset, length int)` — rune-aware substring extraction.
- `CharAt(str string, index int)` — retrieves a rune by position.
- `Excerpt(str, needle string, opts ...ExcerptOption)` — captures text surrounding a needle.
- `Split(str, delimiter string, limit ...int)` — delimiter-based splitting with optional limit.
- `Explode(str string, delimiter string)` — synonym for `Split` without limits.
- `Headline(str string)` — title-cases the string with headline rules.
- `Basename(path string)` / `Dirname(path string)` — filesystem-inspired helpers.
- `Length(str string)` — rune count for the string.
- `WordCount(str string)` — number of words in the string.
- `Words(str string, limit int, end ...string)` — truncates text by word count.

### Transformation & Case

- `ToCamel(str string)` — converts to `camelCase`.
- `ToSnake(str string)` — converts to snake_case using Go rules.
- `Snake(str string, delimiter ...string)` — snake_case with custom delimiter (used by `Kebab`).
- `Kebab(str string)` — outputs kebab-case (`hello-world`).
- `Studly(str string)` — creates StudlyCaps/PascalCase text.
- `Slugify(str string, delimiter rune)` — URL-friendly slugs with Unicode folding.
- `Title(str string)` — title-cases every word.
- `Upper(str string)` / `Lower(str string)` — upper/lower case transforms.
- `UcFirst(str string)` / `LcFirst(str string)` — adjust only the first rune.
- `Start(str, prefix string)` — idempotently prefixes a string.
- `Finish(str, suffix string)` — idempotently suffixes a string.
- `Swap(str string, search string, replace string)` — swaps substrings regardless of order.
- `Mask(str string, mask string, start, length int)` — masks ranges with a repeated mask string.
- `Limit(str string, length int, suffix ...string)` — rune-aware limiter with optional suffix.
- `Truncate(str string, length int, ellipsis string)` — truncates by rune length.
- `NewLine(str string, count ...int)` — appends newline characters (default `\n`).

### Whitespace & Padding

- `Trim(str string, characters ...string)` — trims characters or Unicode whitespace.
- `LTrim(str string, characters ...string)` / `RTrim(str string, characters ...string)` — leading/trailing trims.
- `PadLeft(str, pad string, length int)` / `PadRight(str, pad string, length int)` / `PadBoth(str, pad string, length int)` — symmetrical padding utilities.
- `LeftPad(str string, padStr string, overallLen int)` / `RightPad(str string, padStr string, overallLen int)` — legacy-style padding helpers.
- `Squish(str string)` — collapses consecutive whitespace into single spaces and trims edges.
- `ChopStart(str, prefix string)` / `ChopEnd(str, suffix string)` — remove prefixes/suffixes when present.

### Combination, Replacement & Formatting

- `Append(values ...string)` — concatenates strings efficiently (see also `(*buffer).Append`).
- `Prepend(str string, prefix string)` — adds a prefix once.
- `Repeat(str string, count int)` — repeats the string.
- `Remove(str string, values ...string)` — removes all occurrences of values.
- `RemovePrefix(str string, prefix string)` / `RemoveSuffix(str string, suffix string)` — trim single affixes.
- `Replace(str, search, replace string)` / `ReplaceFirst` / `ReplaceLast` — targeted replacements.
- `ReplaceMatches(str, pattern, replace string)` — regex replacement for each match.
- `ReplaceStart(str, search, replace string)` — prefix-specific replacement.
- `AddSlashes(str string)` — escapes quotes and backslashes.
- `AddSlashesCustom(str string, escapeChars string)` — escapes custom characters.
- `Explode(str string, delimiter string)` — splits into slices (alias of `Split`).

### Encoding, Conversion & Hashing

- `Base64Encode(str string)` / `Base64Decode(str string)` — Base64 conversions.
- `Base32ExtendedEncode(str string)` / `Base32ExtendedDecode(str string)` — Crockford-style Base32 helpers.
- `ToBytes(str string)` — converts to `[]byte` without extra allocation.
- `IntToBase32(n int)` / `IntToBase36(n int)` — integer radix conversions.
- `MD5(str string)` — hex-encoded MD5 hash.
  - ⚠️ Legacy helper; MD5 is insecure and must not be used for security-sensitive hashing.
- `SHA256(str string)` — hex-encoded SHA-256 hash for general-purpose hashing.
- `ToBcryptHash(password string)` — produces a BCrypt hash.
- `BcryptHashCompare(password, hash string)` — verifies BCrypt hashes.

### Generation & Randomness

- `Random(length int)` — cryptographically secure random string of the requested length.
- `RandomFromGamma(length int, shape, scale float64)` — gamma-distributed random strings.

### Price & Numeric Formatting

- `CurrencySymbol(code string, htmlEntity ...bool)` — ISO currency symbol lookup with optional HTML entities.
- `ToPrice(amount float64, code string, htmlEntity ...bool)` — formats prices with symbols.
- `ToPriceFromString(input, code string, htmlEntity ...bool)` — parses string input to a formatted price.
- `ToPriceFromStringOrDefault(input, code, fallback string, htmlEntity ...bool)` — safe parsing with fallback output.
- `Maximum(a, b int)` — max helper commonly used in formatting calculations.

### Flow Control & Conditional Helpers

- `When(str string, condition bool, truthy func(string) string, fallback ...func(string) string)` — conditional string transformer.
- `WhenEmpty`, `WhenNotEmpty`, `WhenContains`, `WhenContainsAll`, `WhenStartsWith`, `WhenEndsWith`, `WhenExactly`, `WhenNotExactly` — call `When` with pre-built predicates.
- `WhenIs`, `WhenIsAscii`, `WhenIsUuid`, `WhenIsUlid` — conditional helpers for pattern/content checks.
- `Unless(str string, predicate func(string) bool, fallback func(string) string)` — inverse of `When` using boolean predicate.
- `Pipe(str string, callback func(string) string)` — transforms via callback returning the new string.
- `Tap(str string, callback func(string))` — observes a string then returns it unchanged.

### Buffer Utility

- `NewBuffer()` — constructs a fluent `*buffer` wrapper around `bytes.Buffer`.
- `(*buffer).Append(v any)` — appends strings, numbers, booleans, runes, or bytes; returns the same buffer for chaining.

### String Builder

- `StringBuilder` — exported struct that accumulates string operations.
- `Of(value string)` — creates a new builder with the given value.
- `(*StringBuilder).Append(parts ...string)` — concatenates pieces.
- `(*StringBuilder).Prepend(prefix string)` — adds a prefix once.
- `(*StringBuilder).LTrim(chars ...string)` / `RTrim(chars ...string)` — trims builder contents in place.
- `(*StringBuilder).Exactly(str string)` — compares against a target value.
- `(*StringBuilder).Unless(predicate func(*StringBuilder) bool, fallback func(*StringBuilder) *StringBuilder)` — conditional mutation.
- `(*StringBuilder).String()` — returns the accumulated string.

### Miscellaneous Utilities

- `Slugify`, `Studly`, `Snake`, `Kebab`, `ToCamel`, `ToSnake`, `Title`, `Upper`, `Lower` — see Transformation section for common naming conversions.
- `Mask`, `Limit`, `Truncate`, `NewLine` — general-purpose helpers noted above but commonly reused across categories.
- `Append`, `Prepend`, `Remove`, `Replace*`, `AddSlashes*` — composition utilities summarized above.

## Best Practices

1. **Use Appropriate Functions**: Choose the most specific function for your task
2. **Handle Errors**: Check for errors when using functions that can fail
3. **Consider Performance**: Some functions may be more efficient than others for your use case
4. **Validate Input**: Always validate input strings before processing them
5. **Use Constants**: Use constants for repeated string values

## License

This package is part of the dracory/base project and is licensed under the same terms. 