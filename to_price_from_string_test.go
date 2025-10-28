package str

import "testing"

func TestToPriceFromString(t *testing.T) {
	tests := []struct {
		name     string
		priceStr string
		currency string
		want     string
		wantErr  bool
	}{
		{
			name:     "Valid price USD",
			priceStr: "19.99",
			currency: "USD",
			want:     "$19.99",
			wantErr:  false,
		},
		{
			name:     "Valid price GBP",
			priceStr: "100.00",
			currency: "GBP",
			want:     "£100.00",
			wantErr:  false,
		},
		{
			name:     "Valid price EUR",
			priceStr: "45.5",
			currency: "EUR",
			want:     "€45.50",
			wantErr:  false,
		},
		{
			name:     "Integer string",
			priceStr: "100",
			currency: "USD",
			want:     "$100.00",
			wantErr:  false,
		},
		{
			name:     "Zero price",
			priceStr: "0",
			currency: "USD",
			want:     "$0.00",
			wantErr:  false,
		},
		{
			name:     "Invalid price - letters",
			priceStr: "abc",
			currency: "USD",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Invalid price - empty",
			priceStr: "",
			currency: "USD",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Invalid price - mixed",
			priceStr: "19.99abc",
			currency: "USD",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToPriceFromString(tt.priceStr, tt.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToPriceFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToPriceFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPriceFromStringOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		priceStr     string
		currency     string
		defaultValue string
		want         string
	}{
		{
			name:         "Valid price",
			priceStr:     "19.99",
			currency:     "USD",
			defaultValue: "n/a",
			want:         "$19.99",
		},
		{
			name:         "Invalid price returns default",
			priceStr:     "invalid",
			currency:     "USD",
			defaultValue: "n/a",
			want:         "n/a",
		},
		{
			name:         "Empty string returns default",
			priceStr:     "",
			currency:     "USD",
			defaultValue: "N/A",
			want:         "N/A",
		},
		{
			name:         "Valid zero price",
			priceStr:     "0",
			currency:     "GBP",
			defaultValue: "n/a",
			want:         "£0.00",
		},
		{
			name:         "Custom default value",
			priceStr:     "abc",
			currency:     "EUR",
			defaultValue: "Price unavailable",
			want:         "Price unavailable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPriceFromStringOrDefault(tt.priceStr, tt.currency, tt.defaultValue); got != tt.want {
				t.Errorf("ToPriceFromStringOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPriceFromStringHTML(t *testing.T) {
	tests := []struct {
		name     string
		priceStr string
		currency string
		want     string
		wantErr  bool
	}{
		{
			name:     "Valid price USD HTML",
			priceStr: "19.99",
			currency: "USD",
			want:     "$19.99",
			wantErr:  false,
		},
		{
			name:     "Valid price GBP HTML",
			priceStr: "100.00",
			currency: "GBP",
			want:     "&pound;100.00",
			wantErr:  false,
		},
		{
			name:     "Valid price EUR HTML",
			priceStr: "45.5",
			currency: "EUR",
			want:     "&euro;45.50",
			wantErr:  false,
		},
		{
			name:     "Valid price JPY HTML",
			priceStr: "1000",
			currency: "JPY",
			want:     "&yen;1000.00",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToPriceFromString(tt.priceStr, tt.currency, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToPriceFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToPriceFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPriceFromStringOrDefaultHTML(t *testing.T) {
	tests := []struct {
		name         string
		priceStr     string
		currency     string
		defaultValue string
		want         string
	}{
		{
			name:         "Valid price GBP HTML",
			priceStr:     "19.99",
			currency:     "GBP",
			defaultValue: "n/a",
			want:         "&pound;19.99",
		},
		{
			name:         "Valid price EUR HTML",
			priceStr:     "50.00",
			currency:     "EUR",
			defaultValue: "n/a",
			want:         "&euro;50.00",
		},
		{
			name:         "Invalid price returns default",
			priceStr:     "invalid",
			currency:     "USD",
			defaultValue: "n/a",
			want:         "n/a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPriceFromStringOrDefault(tt.priceStr, tt.currency, tt.defaultValue, true); got != tt.want {
				t.Errorf("ToPriceFromStringOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
