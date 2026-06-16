package date

import (
	"fmt"
	"testing"
	"time"
)

func d(y int, m time.Month, day int) time.Time {
	return time.Date(y, m, day, 0, 0, 0, 0, time.UTC)
}

func TestAyAdi(t *testing.T) {
	if got := AyAdi(d(2026, time.January, 1)); got != "Ocak" {
		t.Errorf("AyAdi() = %q, want %q", got, "Ocak")
	}
	if got := AyAdi(d(2026, time.December, 1)); got != "Aralık" {
		t.Errorf("AyAdi() = %q, want %q", got, "Aralık")
	}
}

func TestGunAdi(t *testing.T) {
	// 1 January 2026 is a Thursday.
	if got := GunAdi(d(2026, time.January, 1)); got != "Perşembe" {
		t.Errorf("GunAdi() = %q, want %q", got, "Perşembe")
	}
}

func TestFormat(t *testing.T) {
	got := Format(d(2026, time.January, 2))
	want := "2 Ocak 2026 Cuma"
	if got != want {
		t.Errorf("Format() = %q, want %q", got, want)
	}
}

func TestResmiTatil(t *testing.T) {
	tests := []struct {
		date time.Time
		want bool
	}{
		{d(2026, time.January, 1), true},  // Yılbaşı
		{d(2026, time.April, 23), true},   // 23 Nisan
		{d(2026, time.October, 29), true}, // Cumhuriyet
		{d(2026, time.March, 20), true},   // Ramazan Bayramı 1. gün
		{d(2026, time.March, 22), true},   // Ramazan Bayramı 3. gün
		{d(2026, time.March, 23), false},  // bayram bitti
		{d(2020, time.August, 1), true},   // Kurban (ay sınırını aşan)
		{d(2020, time.August, 3), true},   // Kurban son gün
		{d(2026, time.July, 14), false},   // sıradan gün
	}
	for _, tt := range tests {
		if got := ResmiTatil(tt.date); got != tt.want {
			t.Errorf("ResmiTatil(%s) = %v, want %v", tt.date.Format("2006-01-02"), got, tt.want)
		}
	}
}

func TestIsGunu(t *testing.T) {
	// 1 January 2026 (Thursday) is a holiday.
	if IsGunu(d(2026, time.January, 1)) {
		t.Error("Yılbaşı should not be a business day")
	}
	// 3 January 2026 is a Saturday.
	if IsGunu(d(2026, time.January, 3)) {
		t.Error("Saturday should not be a business day")
	}
	// 2 January 2026 (Friday) is an ordinary business day.
	if !IsGunu(d(2026, time.January, 2)) {
		t.Error("2 Jan 2026 should be a business day")
	}
}

func TestIsGunuEkle(t *testing.T) {
	// Friday 2 Jan 2026 + 1 business day -> Monday 5 Jan (skips weekend).
	got := IsGunuEkle(d(2026, time.January, 2), 1)
	if !got.Equal(d(2026, time.January, 5)) {
		t.Errorf("IsGunuEkle(+1) = %s, want 2026-01-05", got.Format("2006-01-02"))
	}
	// Going backward across the weekend.
	got = IsGunuEkle(d(2026, time.January, 5), -1)
	if !got.Equal(d(2026, time.January, 2)) {
		t.Errorf("IsGunuEkle(-1) = %s, want 2026-01-02", got.Format("2006-01-02"))
	}
}

func ExampleFormat() {
	fmt.Println(Format(time.Date(2026, time.January, 2, 0, 0, 0, 0, time.UTC)))
	// Output: 2 Ocak 2026 Cuma
}

func ExampleIsGunuEkle() {
	t := time.Date(2026, time.January, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println(IsGunuEkle(t, 1).Format("2006-01-02"))
	// Output: 2026-01-05
}
