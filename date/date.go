// Package date formats dates with Turkish month and day names and answers
// questions about Turkish public holidays and business days.
package date

import "time"

var months = [...]string{
	"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran",
	"Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık",
}

var days = [...]string{
	"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi",
}

// AyAdi returns the Turkish name of the month, e.g. "Ocak" for January.
func AyAdi(t time.Time) string {
	return months[int(t.Month())-1]
}

// GunAdi returns the Turkish name of the weekday, e.g. "Pazartesi" for Monday.
func GunAdi(t time.Time) string {
	return days[int(t.Weekday())]
}

// Format renders t as "2 Ocak 2006 Pazartesi".
func Format(t time.Time) string {
	return itoa(t.Day()) + " " + AyAdi(t) + " " + itoa(t.Year()) + " " + GunAdi(t)
}

// ResmiTatil reports whether t falls on a Turkish public holiday. Fixed
// national holidays are computed; religious holidays (Ramazan and Kurban
// Bayramı) follow the lunar calendar and are read from an embedded table that
// currently covers 2020-2035.
func ResmiTatil(t time.Time) bool {
	_, m, d := t.Date()
	for _, h := range fixedHolidays {
		if h.month == m && h.day == d {
			return true
		}
	}
	day := truncate(t)
	for _, b := range religiousHolidays {
		start := time.Date(b.year, b.month, b.day, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 0, b.days)
		if !day.Before(start) && day.Before(end) {
			return true
		}
	}
	return false
}

// truncate returns the calendar date of t in UTC, stripping the clock time and
// location so that date comparisons are not skewed by time zones.
func truncate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// IsGunu reports whether t is a business day: a weekday that is not a public
// holiday.
func IsGunu(t time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	}
	return !ResmiTatil(t)
}

// IsGunuEkle returns the date n business days from t. A positive n moves
// forward, a negative n moves backward. The starting day is never counted.
func IsGunuEkle(t time.Time, n int) time.Time {
	step := 1
	if n < 0 {
		step = -1
		n = -n
	}
	for n > 0 {
		t = t.AddDate(0, 0, step)
		if IsGunu(t) {
			n--
		}
	}
	return t
}

type fixedHoliday struct {
	month time.Month
	day   int
}

// fixedHolidays lists the national holidays that fall on the same date every
// year.
var fixedHolidays = []fixedHoliday{
	{time.January, 1},  // Yılbaşı
	{time.April, 23},   // Ulusal Egemenlik ve Çocuk Bayramı
	{time.May, 1},      // Emek ve Dayanışma Günü
	{time.May, 19},     // Atatürk'ü Anma, Gençlik ve Spor Bayramı
	{time.July, 15},    // Demokrasi ve Millî Birlik Günü
	{time.August, 30},  // Zafer Bayramı
	{time.October, 29}, // Cumhuriyet Bayramı
}

// bayram is the first day of a religious holiday and how many full days it
// lasts (3 for Ramazan, 4 for Kurban).
type bayram struct {
	year  int
	month time.Month
	day   int
	days  int
}

// religiousHolidays lists the lunar-calendar holidays for 2020-2035. The dates
// for 2020-2026 are the official observed dates; 2027 onward follow Diyanet's
// pre-calculated calendar and may shift by a day in rare cases.
var religiousHolidays = []bayram{
	// Ramazan Bayramı (3 days)
	{2020, time.May, 24, 3},
	{2021, time.May, 13, 3},
	{2022, time.May, 2, 3},
	{2023, time.April, 21, 3},
	{2024, time.April, 10, 3},
	{2025, time.March, 30, 3},
	{2026, time.March, 20, 3},
	{2027, time.March, 9, 3},
	{2028, time.February, 26, 3},
	{2029, time.February, 14, 3},
	{2030, time.February, 4, 3},
	{2031, time.January, 24, 3},
	{2032, time.January, 14, 3},
	{2033, time.January, 2, 3},
	{2034, time.December, 12, 3},
	{2035, time.December, 1, 3},
	// Kurban Bayramı (4 days)
	{2020, time.July, 31, 4},
	{2021, time.July, 20, 4},
	{2022, time.July, 9, 4},
	{2023, time.June, 28, 4},
	{2024, time.June, 16, 4},
	{2025, time.June, 6, 4},
	{2026, time.May, 27, 4},
	{2027, time.May, 16, 4},
	{2028, time.May, 5, 4},
	{2029, time.April, 24, 4},
	{2030, time.April, 13, 4},
	{2031, time.April, 2, 4},
	{2032, time.March, 22, 4},
	{2033, time.March, 11, 4},
	{2034, time.March, 1, 4},
	{2035, time.February, 18, 4},
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
