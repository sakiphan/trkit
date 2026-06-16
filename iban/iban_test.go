package iban

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"TR330006100519786457841326", true},
		{"TR33 0006 1005 1978 6457 8413 26", true},
		{"tr330006100519786457841326", true},
		{"TR340006100519786457841326", false},  // bad checksum
		{"DE89370400440532013000", false},      // wrong country
		{"TR33000610051978645784132", false},   // too short
		{"TR3300061005197864578413266", false}, // too long
		{"TR3300061005197864578413X6", false},  // non-digit body
		{"", false},
	}
	for _, tt := range tests {
		if got := Valid(tt.in); got != tt.want {
			t.Errorf("Valid(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestFormat(t *testing.T) {
	got := Format("TR330006100519786457841326")
	want := "TR33 0006 1005 1978 6457 8413 26"
	if got != want {
		t.Errorf("Format() = %q, want %q", got, want)
	}
}

func TestMask(t *testing.T) {
	got := Mask("TR330006100519786457841326")
	want := "TR33 **** **** **** **** **13 26"
	if got != want {
		t.Errorf("Mask() = %q, want %q", got, want)
	}
}

func TestBankaKodu(t *testing.T) {
	got, err := BankaKodu("TR330006100519786457841326")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "00061" {
		t.Errorf("BankaKodu() = %q, want %q", got, "00061")
	}
	if _, err := BankaKodu("TR33"); err == nil {
		t.Error("expected error for short IBAN")
	}
}

func ExampleValid() {
	fmt.Println(Valid("TR33 0006 1005 1978 6457 8413 26"))
	// Output: true
}

func ExampleMask() {
	fmt.Println(Mask("TR330006100519786457841326"))
	// Output: TR33 **** **** **** **** **13 26
}
