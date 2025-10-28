package str

import "strings"

// currencyInfo holds the symbol and HTML entity for a currency
type currencyInfo struct {
	code       string
	symbol     string
	htmlEntity string
}

// currencyMap contains all supported currencies
var currencyMap = map[string]currencyInfo{
	// Major currencies
	"USD": {"USD", "$", "$"},
	"EUR": {"EUR", "€", "&euro;"},
	"GBP": {"GBP", "£", "&pound;"},
	"JPY": {"JPY", "¥", "&yen;"},
	"CNY": {"CNY", "¥", "&yen;"},
	"INR": {"INR", "₹", "&#8377;"},
	"KRW": {"KRW", "₩", "&#8361;"},
	"RUB": {"RUB", "₽", "&#8381;"},
	"TRY": {"TRY", "₺", "&#8378;"},

	// Dollar variants
	"AUD": {"AUD", "A$", "A$"},
	"CAD": {"CAD", "C$", "C$"},
	"HKD": {"HKD", "HK$", "HK$"},
	"SGD": {"SGD", "S$", "S$"},
	"NZD": {"NZD", "NZ$", "NZ$"},
	"MXN": {"MXN", "Mex$", "Mex$"},
	"BRL": {"BRL", "R$", "R$"},
	"ARS": {"ARS", "$", "$"},
	"CLP": {"CLP", "$", "$"},
	"COP": {"COP", "$", "$"},

	// European currencies
	"CHF": {"CHF", "Fr", "Fr"},
	"SEK": {"SEK", "kr", "kr"},
	"NOK": {"NOK", "kr", "kr"},
	"DKK": {"DKK", "kr", "kr"},
	"PLN": {"PLN", "zł", "z&#322;"},
	"CZK": {"CZK", "Kč", "K&#269;"},
	"HUF": {"HUF", "Ft", "Ft"},
	"RON": {"RON", "lei", "lei"},
	"BGN": {"BGN", "лв", "&#1083;&#1074;"},

	// Middle East & Africa
	"SAR": {"SAR", "﷼", "&#65020;"},
	"AED": {"AED", "د.إ", "&#1583;.&#1573;"},
	"ILS": {"ILS", "₪", "&#8362;"},
	"ZAR": {"ZAR", "R", "R"},
	"EGP": {"EGP", "£", "&pound;"},
	"NGN": {"NGN", "₦", "&#8358;"},
	"KES": {"KES", "KSh", "KSh"},

	// Asian currencies
	"THB": {"THB", "฿", "&#3647;"},
	"IDR": {"IDR", "Rp", "Rp"},
	"MYR": {"MYR", "RM", "RM"},
	"PHP": {"PHP", "₱", "&#8369;"},
	"VND": {"VND", "₫", "&#8363;"},
	"PKR": {"PKR", "₨", "&#8360;"},
	"BDT": {"BDT", "৳", "&#2547;"},

	// Cryptocurrencies
	"BTC": {"BTC", "₿", "&#8383;"},
	"ETH": {"ETH", "Ξ", "&#926;"},
}

// CurrencySymbol returns the symbol for the given ISO 4217 currency code.
// By default, returns the Unicode symbol. Pass true to get HTML entity instead.
// Returns the original currency code if not recognized.
// Example: CurrencySymbol("USD") returns "$"
// Example: CurrencySymbol("GBP", true) returns "&pound;"
func CurrencySymbol(currencyCode string, htmlEntity ...bool) string {
	useHTML := false
	if len(htmlEntity) > 0 {
		useHTML = htmlEntity[0]
	}

	info, exists := currencyMap[strings.ToUpper(currencyCode)]
	if !exists {
		return currencyCode
	}

	if useHTML {
		return info.htmlEntity
	}
	return info.symbol
}
