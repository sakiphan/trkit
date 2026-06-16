// Package money formats Turkish Lira amounts and spells them out in words.
package money

import (
	"strconv"
	"strings"
)

// FormatKurus formats an amount given in kuruş (1 lira = 100 kuruş) using the
// Turkish convention: '.' as the thousands separator, ',' as the decimal
// separator, and the ₺ sign. It works on integers to avoid float rounding.
//
//	FormatKurus(123456789) == "1.234.567,89 ₺"
func FormatKurus(kurus int64) string {
	neg := kurus < 0
	if neg {
		kurus = -kurus
	}
	lira := kurus / 100
	rem := kurus % 100

	s := groupThousands(strconv.FormatInt(lira, 10)) +
		"," + twoDigits(rem) + " ₺"
	if neg {
		s = "-" + s
	}
	return s
}

// Format is a convenience wrapper that formats a lira amount expressed as a
// float. The value is rounded to the nearest kuruş.
func Format(f float64) string {
	return FormatKurus(int64(roundHalfAway(f * 100)))
}

func roundHalfAway(f float64) float64 {
	if f < 0 {
		return -roundHalfAway(-f)
	}
	return float64(int64(f + 0.5))
}

func twoDigits(n int64) string {
	if n < 10 {
		return "0" + strconv.FormatInt(n, 10)
	}
	return strconv.FormatInt(n, 10)
}

var ones = [...]string{"", "bir", "iki", "üç", "dört", "beş", "altı", "yedi", "sekiz", "dokuz"}
var tens = [...]string{"", "on", "yirmi", "otuz", "kırk", "elli", "altmış", "yetmiş", "seksen", "doksan"}
var scales = [...]string{"", "bin", "milyon", "milyar", "trilyon", "katrilyon", "kentilyon"}

// Yaziyla spells out a lira amount in Turkish words, e.g. 1234 becomes
// "bin iki yüz otuz dört lira". Note the Turkish convention of omitting the
// leading "bir" for exactly one thousand ("bin", not "bir bin").
func Yaziyla(lira int64) string {
	if lira == 0 {
		return "sıfır lira"
	}
	neg := lira < 0
	if neg {
		lira = -lira
	}

	var groups []int64
	for lira > 0 {
		groups = append(groups, lira%1000)
		lira /= 1000
	}

	var parts []string
	for i := len(groups) - 1; i >= 0; i-- {
		g := groups[i]
		if g == 0 {
			continue
		}
		if i == 1 && g == 1 {
			parts = append(parts, "bin")
			continue
		}
		parts = append(parts, threeWords(g))
		if scales[i] != "" {
			parts = append(parts, scales[i])
		}
	}

	out := strings.Join(parts, " ")
	if neg {
		out = "eksi " + out
	}
	return out + " lira"
}

// threeWords spells out a number in the range 1..999.
func threeWords(n int64) string {
	var parts []string
	h := n / 100
	t := (n % 100) / 10
	o := n % 10
	if h == 1 {
		parts = append(parts, "yüz")
	} else if h > 1 {
		parts = append(parts, ones[h], "yüz")
	}
	if t > 0 {
		parts = append(parts, tens[t])
	}
	if o > 0 {
		parts = append(parts, ones[o])
	}
	return strings.Join(parts, " ")
}

// groupThousands inserts '.' every three digits from the right.
func groupThousands(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var b strings.Builder
	first := n % 3
	if first == 0 {
		first = 3
	}
	b.WriteString(s[:first])
	for i := first; i < n; i += 3 {
		b.WriteByte('.')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
