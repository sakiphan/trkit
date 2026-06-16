package identity

import (
	"fmt"
	"testing"
)

func TestValidTCKN(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"10000000832", true},
		{"19191919190", true},
		{"11111111110", true},
		{"00000000000", false},  // first digit 0
		{"12345678901", false},  // bad check digit
		{"1000000083", false},   // 10 digits
		{"100000008320", false}, // 12 digits
		{"1000000083a", false},  // non-digit
		{"", false},
	}
	for _, tt := range tests {
		if got := ValidTCKN(tt.in); got != tt.want {
			t.Errorf("ValidTCKN(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestValidVKN(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"1234567899", true},
		{"4540536921", true},
		{"1234567890", false},
		{"123456789", false},   // 9 digits
		{"12345678990", false}, // 11 digits
		{"123456789a", false},  // non-digit
		{"", false},
	}
	for _, tt := range tests {
		if got := ValidVKN(tt.in); got != tt.want {
			t.Errorf("ValidVKN(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func ExampleValidTCKN() {
	fmt.Println(ValidTCKN("10000000832"))
	fmt.Println(ValidTCKN("12345678901"))
	// Output:
	// true
	// false
}

func ExampleValidVKN() {
	fmt.Println(ValidVKN("1234567899"))
	// Output:
	// true
}
