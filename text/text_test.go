package text

import (
	"fmt"
	"testing"
)

func TestUpper(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"istanbul", "İSTANBUL"},
		{"ığdır", "IĞDIR"},
		{"diyarbakır", "DİYARBAKIR"},
		{"çağrı", "ÇAĞRI"},
		{"şişli", "ŞİŞLİ"},
		{"abc", "ABC"},
	}
	for _, tt := range tests {
		if got := Upper(tt.in); got != tt.want {
			t.Errorf("Upper(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"İSTANBUL", "istanbul"},
		{"DİYARBAKIR", "diyarbakır"},
		{"IĞDIR", "ığdır"},
		{"ÇAĞRI", "çağrı"},
		{"ABC", "abc"},
	}
	for _, tt := range tests {
		if got := Lower(tt.in); got != tt.want {
			t.Errorf("Lower(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestTitle(t *testing.T) {
	got := Title("istanbul büyükşehir belediyesi")
	want := "İstanbul Büyükşehir Belediyesi"
	if got != want {
		t.Errorf("Title() = %q, want %q", got, want)
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"ığdır", "Iğdır"},
		{"istanbul", "İstanbul"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := Capitalize(tt.in); got != tt.want {
			t.Errorf("Capitalize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestSlug(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Çağrı'nın İşi", "cagrinin-isi"},
		{"İstanbul Büyükşehir", "istanbul-buyuksehir"},
		{"  boşluk   testi  ", "bosluk-testi"},
		{"Merhaba, Dünya!", "merhaba-dunya"},
		{"ÜÇ- --AĞAÇ", "uc-agac"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := Slug(tt.in); got != tt.want {
			t.Errorf("Slug(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func ExampleUpper() {
	fmt.Println(Upper("istanbul"))
	// Output: İSTANBUL
}

func ExampleLower() {
	fmt.Println(Lower("DİYARBAKIR"))
	// Output: diyarbakır
}

func ExampleSlug() {
	fmt.Println(Slug("Çağrı'nın İşi"))
	// Output: cagrinin-isi
}
