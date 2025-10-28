package str

import "strconv"

// ToPriceFromString converts a string price to a formatted price string with the given ISO 4217 currency code.
// Returns an error if the string cannot be parsed as a float.
// The price is formatted to 2 decimal places.
// By default, returns Unicode symbol. Pass true to get HTML entity instead.
// Example: ToPriceFromString("19.99", "USD") returns "$19.99", nil
// Example: ToPriceFromString("19.99", "GBP", true) returns "&pound;19.99", nil
func ToPriceFromString(priceStr string, currencyCode string, htmlEntity ...bool) (string, error) {
	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return "", err
	}
	return ToPrice(priceFloat, currencyCode, htmlEntity...), nil
}

// ToPriceFromStringOrDefault converts a string price to a formatted price string with the given ISO 4217 currency code.
// Returns the defaultValue if the string cannot be parsed as a float.
// The price is formatted to 2 decimal places.
// By default, returns Unicode symbol. Pass true to get HTML entity instead.
// Example: ToPriceFromStringOrDefault("19.99", "USD", "n/a") returns "$19.99"
// Example: ToPriceFromStringOrDefault("invalid", "USD", "n/a") returns "n/a"
// Example: ToPriceFromStringOrDefault("19.99", "GBP", "n/a", true) returns "&pound;19.99"
func ToPriceFromStringOrDefault(priceStr string, currencyCode string, defaultValue string, htmlEntity ...bool) string {
	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return defaultValue
	}
	return ToPrice(priceFloat, currencyCode, htmlEntity...)
}
