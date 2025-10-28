package str

import "strconv"

// ToPrice converts a float64 price to a formatted string with the given ISO 4217 currency code.
// The price is formatted to 2 decimal places.
// By default, returns Unicode symbol. Pass true to get HTML entity instead.
// Example: ToPrice(19.99, "USD") returns "$19.99"
// Example: ToPrice(19.99, "GBP", true) returns "&pound;19.99"
func ToPrice(price float64, currencyCode string, htmlEntity ...bool) string {
	priceStr := strconv.FormatFloat(price, 'f', 2, 64)
	return CurrencySymbol(currencyCode, htmlEntity...) + priceStr
}
