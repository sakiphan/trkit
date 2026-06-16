// Package phone normalizes, validates and formats Turkish mobile numbers.
package phone

import (
	"errors"
	"strings"
)

// ErrInvalid is returned when a number cannot be parsed as a Turkish mobile
// number.
var ErrInvalid = errors.New("phone: invalid number")

// Normalize parses any common representation of a Turkish mobile number and
// returns it in E.164 form, e.g. "+905321234567". Accepted inputs include
// "0532 123 45 67", "+90 532 123 45 67", "905321234567" and "5321234567".
func Normalize(s string) (string, error) {
	d := digits(s)

	switch {
	case strings.HasPrefix(d, "90") && len(d) == 12:
		d = d[2:]
	case strings.HasPrefix(d, "0") && len(d) == 11:
		d = d[1:]
	}

	if len(d) != 10 || d[0] != '5' {
		return "", ErrInvalid
	}
	return "+90" + d, nil
}

// Valid reports whether s is a valid Turkish mobile number.
func Valid(s string) bool {
	_, err := Normalize(s)
	return err == nil
}

// Format returns the national display form of a number: "0532 123 45 67".
// It returns the original string unchanged if it cannot be parsed.
func Format(s string) string {
	n, err := Normalize(s)
	if err != nil {
		return s
	}
	d := n[3:] // strip "+90"
	return "0" + d[0:3] + " " + d[3:6] + " " + d[6:8] + " " + d[8:10]
}

// Operator returns a best-effort carrier name based on the number prefix.
// Because of mobile number portability the result is not guaranteed; an empty
// string is returned for unknown or invalid prefixes.
func Operator(s string) string {
	n, err := Normalize(s)
	if err != nil {
		return ""
	}
	// p is the three-digit prefix, e.g. "532". Allocations are made in
	// contiguous blocks, so a range check is enough.
	p := n[3:6]
	switch {
	case p >= "500" && p <= "509":
		return "Türk Telekom"
	case p >= "530" && p <= "539":
		return "Turkcell"
	case p >= "540" && p <= "549":
		return "Vodafone"
	case p >= "550" && p <= "559":
		return "Türk Telekom"
	default:
		return ""
	}
}

func digits(s string) string {
	var b strings.Builder
	for _, c := range s {
		if c >= '0' && c <= '9' {
			b.WriteRune(c)
		}
	}
	return b.String()
}
