// Package iban validates and formats Turkish IBANs.
package iban

import "strings"

// length of a Turkish IBAN without spaces.
const trLen = 26

// Valid reports whether s is a valid Turkish IBAN. Spaces are ignored.
// It checks the country code, length and the ISO 7064 mod-97 checksum.
func Valid(s string) bool {
	s = clean(s)
	if len(s) != trLen {
		return false
	}
	if s[0] != 'T' || s[1] != 'R' {
		return false
	}
	for i := 2; i < trLen; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return mod97(s) == 1
}

// Format groups an IBAN into blocks of four for display:
// "TR33 0006 1005 1978 6457 8413 26".
func Format(s string) string {
	return group(clean(s))
}

// Mask hides the middle of an IBAN, keeping the first and last digits visible:
// "TR33 **** **** **** **** **13 26".
func Mask(s string) string {
	s = clean(s)
	if len(s) != trLen {
		return Format(s)
	}
	b := []byte(s)
	for i := 4; i < trLen-4; i++ {
		b[i] = '*'
	}
	return group(string(b))
}

// BankaKodu returns the 5-digit bank code (positions 5-9) of the IBAN.
func BankaKodu(s string) (string, error) {
	s = clean(s)
	if len(s) != trLen {
		return "", ErrInvalid
	}
	return s[4:9], nil
}

// ErrInvalid is returned when an IBAN does not have the expected shape.
var ErrInvalid = errInvalid{}

type errInvalid struct{}

func (errInvalid) Error() string { return "iban: invalid IBAN" }

func clean(s string) string {
	s = strings.ToUpper(strings.TrimSpace(s))
	return strings.ReplaceAll(s, " ", "")
}

func group(s string) string {
	var b strings.Builder
	for i, c := range s {
		if i > 0 && i%4 == 0 {
			b.WriteByte(' ')
		}
		b.WriteRune(c)
	}
	return b.String()
}

// mod97 computes the ISO 7064 mod-97 checksum without big integers. The first
// four characters are moved to the end, letters are expanded to two-digit
// numbers (A=10..Z=35) and the remainder is accumulated chunk by chunk.
func mod97(s string) int {
	rearranged := s[4:] + s[:4]
	rem := 0
	for _, c := range rearranged {
		switch {
		case c >= '0' && c <= '9':
			rem = (rem*10 + int(c-'0')) % 97
		case c >= 'A' && c <= 'Z':
			v := int(c-'A') + 10
			rem = (rem*100 + v) % 97
		default:
			return -1
		}
	}
	return rem
}
