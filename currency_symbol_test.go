package str

import (
	"strings"
	"testing"
)

func TestCurrencySymbol(t *testing.T) {
	// Test all currencies from the map - Unicode symbols
	for code, info := range currencyMap {
		t.Run(code+" Unicode", func(t *testing.T) {
			got := CurrencySymbol(code)
			if got != info.symbol {
				t.Errorf("CurrencySymbol(%q) = %q, want %q", code, got, info.symbol)
			}
		})

		// Test lowercase
		t.Run(code+" lowercase", func(t *testing.T) {
			got := CurrencySymbol(strings.ToLower(code))
			if got != info.symbol {
				t.Errorf("CurrencySymbol(%q) = %q, want %q", strings.ToLower(code), got, info.symbol)
			}
		})
	}

	// Test unknown currency
	t.Run("Unknown currency", func(t *testing.T) {
		unknown := "XXX"
		got := CurrencySymbol(unknown)
		if got != unknown {
			t.Errorf("CurrencySymbol(%q) = %q, want %q", unknown, got, unknown)
		}
	})

	// Test empty string
	t.Run("Empty string", func(t *testing.T) {
		got := CurrencySymbol("")
		if got != "" {
			t.Errorf("CurrencySymbol(\"\") = %q, want \"\"", got)
		}
	})
}

func TestCurrencySymbolHTML(t *testing.T) {
	// Test all currencies from the map - HTML entities
	for code, info := range currencyMap {
		t.Run(code+" HTML", func(t *testing.T) {
			got := CurrencySymbol(code, true)
			if got != info.htmlEntity {
				t.Errorf("CurrencySymbol(%q, true) = %q, want %q", code, got, info.htmlEntity)
			}
		})

		// Test lowercase with HTML
		t.Run(code+" lowercase HTML", func(t *testing.T) {
			got := CurrencySymbol(strings.ToLower(code), true)
			if got != info.htmlEntity {
				t.Errorf("CurrencySymbol(%q, true) = %q, want %q", strings.ToLower(code), got, info.htmlEntity)
			}
		})
	}

	// Test unknown currency with HTML
	t.Run("Unknown currency HTML", func(t *testing.T) {
		unknown := "XXX"
		got := CurrencySymbol(unknown, true)
		if got != unknown {
			t.Errorf("CurrencySymbol(%q, true) = %q, want %q", unknown, got, unknown)
		}
	})
}
