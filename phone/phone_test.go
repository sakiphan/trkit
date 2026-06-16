package phone

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T) {
	want := "+905321234567"
	inputs := []string{
		"05321234567",
		"5321234567",
		"905321234567",
		"+905321234567",
		"0532 123 45 67",
		"+90 532 123 45 67",
		"(0532) 123 45 67",
	}
	for _, in := range inputs {
		got, err := Normalize(in)
		if err != nil {
			t.Errorf("Normalize(%q) returned error: %v", in, err)
			continue
		}
		if got != want {
			t.Errorf("Normalize(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestNormalizeInvalid(t *testing.T) {
	inputs := []string{
		"",
		"0212 123 45 67", // landline
		"4321234567",     // does not start with 5
		"053212345",      // too short
		"053212345678",   // too long
	}
	for _, in := range inputs {
		if _, err := Normalize(in); err == nil {
			t.Errorf("Normalize(%q) expected error", in)
		}
	}
}

func TestFormat(t *testing.T) {
	got := Format("+905321234567")
	if got != "0532 123 45 67" {
		t.Errorf("Format() = %q, want %q", got, "0532 123 45 67")
	}
	if got := Format("not a number"); got != "not a number" {
		t.Errorf("Format() should pass through invalid input, got %q", got)
	}
}

func TestOperator(t *testing.T) {
	tests := map[string]string{
		"05321234567": "Turkcell",
		"05441234567": "Vodafone",
		"05011234567": "Türk Telekom",
		"05061234567": "Türk Telekom",
		"05511234567": "Türk Telekom",
		"05591234567": "Türk Telekom",
	}
	for in, want := range tests {
		if got := Operator(in); got != want {
			t.Errorf("Operator(%q) = %q, want %q", in, got, want)
		}
	}
	if got := Operator("not a number"); got != "" {
		t.Errorf("Operator() = %q, want empty", got)
	}
}

func ExampleNormalize() {
	n, _ := Normalize("0532 123 45 67")
	fmt.Println(n)
	// Output: +905321234567
}

func ExampleFormat() {
	fmt.Println(Format("+905321234567"))
	// Output: 0532 123 45 67
}
