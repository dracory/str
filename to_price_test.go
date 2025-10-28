package str

import "testing"

func TestToPrice(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		currency string
		want     string
	}{
		{
			name:     "USD with decimal",
			price:    19.99,
			currency: "USD",
			want:     "$19.99",
		},
		{
			name:     "GBP whole number",
			price:    100.00,
			currency: "GBP",
			want:     "£100.00",
		},
		{
			name:     "EUR with cents",
			price:    45.50,
			currency: "EUR",
			want:     "€45.50",
		},
		{
			name:     "Zero price",
			price:    0.00,
			currency: "USD",
			want:     "$0.00",
		},
		{
			name:     "Large price",
			price:    1234567.89,
			currency: "GBP",
			want:     "£1234567.89",
		},
		{
			name:     "JPY currency",
			price:    99.99,
			currency: "JPY",
			want:     "¥99.99",
		},
		{
			name:     "Lowercase currency",
			price:    25.00,
			currency: "usd",
			want:     "$25.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPrice(tt.price, tt.currency); got != tt.want {
				t.Errorf("ToPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPriceHTML(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		currency string
		want     string
	}{
		{
			name:     "USD with HTML",
			price:    19.99,
			currency: "USD",
			want:     "$19.99",
		},
		{
			name:     "GBP with HTML",
			price:    100.00,
			currency: "GBP",
			want:     "&pound;100.00",
		},
		{
			name:     "EUR with HTML",
			price:    45.50,
			currency: "EUR",
			want:     "&euro;45.50",
		},
		{
			name:     "JPY with HTML",
			price:    99.99,
			currency: "JPY",
			want:     "&yen;99.99",
		},
		{
			name:     "INR with HTML",
			price:    1500.00,
			currency: "INR",
			want:     "&#8377;1500.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPrice(tt.price, tt.currency, true); got != tt.want {
				t.Errorf("ToPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
