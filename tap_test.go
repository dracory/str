package str

import "testing"

func TestTap(t *testing.T) {
	t.Run("nil callback returns original", func(t *testing.T) {
		if got := Tap("hello", nil); got != "hello" {
			t.Fatalf("Tap() = %q, want %q", got, "hello")
		}
	})

	t.Run("callback receives string", func(t *testing.T) {
		var received string
		result := Tap("value", func(s string) {
			received = s
		})

		if result != "value" {
			t.Fatalf("Tap() returned %q, want %q", result, "value")
		}
		if received != "value" {
			t.Fatalf("callback received %q, want %q", received, "value")
		}
	})
}
