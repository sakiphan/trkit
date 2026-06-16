package money

import (
	"fmt"
	"testing"
)

func TestFormatKurus(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "0,00 ₺"},
		{5, "0,05 ₺"},
		{99, "0,99 ₺"},
		{100, "1,00 ₺"},
		{123456789, "1.234.567,89 ₺"},
		{100000000, "1.000.000,00 ₺"},
		{-12345, "-123,45 ₺"},
	}
	for _, tt := range tests {
		if got := FormatKurus(tt.in); got != tt.want {
			t.Errorf("FormatKurus(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFormat(t *testing.T) {
	if got := Format(1234.5); got != "1.234,50 ₺" {
		t.Errorf("Format(1234.5) = %q, want %q", got, "1.234,50 ₺")
	}
	if got := Format(0.1); got != "0,10 ₺" {
		t.Errorf("Format(0.1) = %q, want %q", got, "0,10 ₺")
	}
}

func TestYaziyla(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "sıfır lira"},
		{1, "bir lira"},
		{11, "on bir lira"},
		{100, "yüz lira"},
		{1000, "bin lira"},
		{1234, "bin iki yüz otuz dört lira"},
		{2000, "iki bin lira"},
		{1000000, "bir milyon lira"},
		{1234567, "bir milyon iki yüz otuz dört bin beş yüz altmış yedi lira"},
		{-50, "eksi elli lira"},
	}
	for _, tt := range tests {
		if got := Yaziyla(tt.in); got != tt.want {
			t.Errorf("Yaziyla(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func ExampleFormatKurus() {
	fmt.Println(FormatKurus(123456789))
	// Output: 1.234.567,89 ₺
}

func ExampleYaziyla() {
	fmt.Println(Yaziyla(1234))
	// Output: bin iki yüz otuz dört lira
}
